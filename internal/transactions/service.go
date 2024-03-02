package transactions

import (
	"context"
	"database/sql"

	"github.com/anti-duhring/goncurrency/internal/clients"
)

type Service struct {
	Repository        Repository
	ClientsRepository clients.Repository
}

func NewService(repository Repository, clientsRepository clients.Repository) *Service {
	return &Service{Repository: repository, ClientsRepository: clientsRepository}
}

func (s *Service) GetTransactionsFromClient(ctx context.Context, id int, limit int) (*[]Transaction, error) {
	transactions, err := s.Repository.FindManyByClientID(ctx, id, limit)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *Service) CreateTransaction(ctx context.Context, clientId int, t *Transaction) (*clients.ClientExtract, error) {
	clientExtract := &clients.ClientExtract{}
	err := s.Repository.WithTransaction(ctx, func(tx *sql.Tx) error {
		client, err := s.ClientsRepository.FindOneByID(ctx, clientId, tx)
		if err != nil {
			tx.Rollback()
			return err
		}

		var newBalance int

		if t.Operation == "d" {
			newBalance = client.Balance - t.Amount
		}
		if t.Operation == "c" {
			newBalance = client.Balance + t.Amount
		}

		if newBalance < client.AccountLimit*-1 {
			tx.Rollback()
			return ErrAccountLimitExceeded
		}

		inserOneInput := InsertOneInput{
			ID:          clientId,
			Transaction: t,
		}
		if err = s.Repository.InsertOne(ctx, inserOneInput, tx); err != nil {
			tx.Rollback()
			return err
		}

		updateBalanceInput := clients.UpdateBalanceByIDInput{
			ID:      clientId,
			Balance: newBalance,
		}
		if err = s.ClientsRepository.UpdateBalanceByID(ctx, updateBalanceInput, tx); err != nil {
			tx.Rollback()
			return err
		}

		clientExtract.AccountLimit = client.AccountLimit
		clientExtract.Balance = newBalance

		return nil
	})
	if err != nil {
		return nil, err
	}

	return clientExtract, nil
}
