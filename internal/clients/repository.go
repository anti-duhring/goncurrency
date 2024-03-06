package clients

import (
	"context"
	"database/sql"
	"errors"
)

type Repository interface {
	FindOneByID(ctx context.Context, id int, tx *sql.Tx) (*Client, error)
	UpdateBalanceByID(ctx context.Context, input UpdateBalanceByIDInput, tx *sql.Tx) error
}

type RepositoryPostgres struct {
	DB *sql.DB
}

func NewRepositoryPostgres(db *sql.DB) *RepositoryPostgres {
	return &RepositoryPostgres{DB: db}
}

func (r *RepositoryPostgres) FindOneByID(ctx context.Context, id int, tx *sql.Tx) (*Client, error) {
	var client Client
	var row *sql.Row

	query := "SELECT account_limit, balance FROM clients WHERE id = $1 FOR UPDATE"

	if tx != nil {
		row = tx.QueryRowContext(ctx, query, id)
	} else {
		row = r.DB.QueryRowContext(ctx, query, id)
	}

	err := row.Scan(&client.AccountLimit, &client.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrClientNotFound
		}
		return nil, err
	}

	return &client, nil
}

type UpdateBalanceByIDInput struct {
	ID      int
	Balance int
}

func (r *RepositoryPostgres) UpdateBalanceByID(ctx context.Context, input UpdateBalanceByIDInput, tx *sql.Tx) error {
	query := "UPDATE clients SET balance = $1 WHERE id = $2"

	if tx != nil {
		_, err := tx.ExecContext(ctx, query, input.Balance, input.ID)
		if err != nil {
			return err
		}

		return nil
	}

	_, err := r.DB.ExecContext(ctx, query, input.Balance, input.ID)
	if err != nil {
		return err
	}

	return nil
}
