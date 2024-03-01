package transactions

import (
	"context"
)

type Service struct {
	Repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{Repository: repository}
}

func (s *Service) GetTransactionsFromClient(ctx context.Context, id int, limit int) (*[]Transaction, error) {
	transactions, err := s.Repository.FindManyByClientID(ctx, id, limit)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *Service) CreateTransaction(ctx context.Context, id int, t *Transaction) error {
	if err := s.Repository.InsertOne(ctx, id, t); err != nil {
		return err
	}

	return nil
}
