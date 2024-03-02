package transactions

import (
	"context"
	"database/sql"
)

type Repository interface {
	FindManyByClientID(ctx context.Context, id int, limit int) (*[]Transaction, error)
	InsertOne(ctx context.Context, input InsertOneInput, tx *sql.Tx) error
	WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error
}

type RepositoryPostgres struct {
	DB *sql.DB
}

func NewRepositoryPostgres(db *sql.DB) *RepositoryPostgres {
	return &RepositoryPostgres{DB: db}
}

func (r *RepositoryPostgres) FindManyByClientID(ctx context.Context, id int, limit int) (*[]Transaction, error) {
	row, err := r.DB.QueryContext(ctx, "SELECT amount, operation, description, created_at FROM transactions WHERE client_id = $1 ORDER BY created_at DESC LIMIT $2", id, limit)
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	for row.Next() {
		var t Transaction
		if err := row.Scan(&t.Amount, &t.Operation, &t.Description, &t.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return &transactions, nil
}

type InsertOneInput struct {
	ID          int
	Transaction *Transaction
}

func (r *RepositoryPostgres) InsertOne(ctx context.Context, input InsertOneInput, tx *sql.Tx) error {
	query := "INSERT INTO transactions (client_id, amount, operation, description) VALUES ($1, $2, $3, $4)"

	if tx != nil {
		_, err := tx.ExecContext(ctx, query, input.ID, input.Transaction.Amount, input.Transaction.Operation, input.Transaction.Description)
		return err
	}

	_, err := r.DB.ExecContext(ctx, query, input.ID, input.Transaction.Amount, input.Transaction.Operation, input.Transaction.Description)
	return err
}

func (r *RepositoryPostgres) WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}
