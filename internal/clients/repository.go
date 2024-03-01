package clients

import (
	"context"
	"database/sql"
	"errors"
)

type Repository interface {
	FindOneByID(ctx context.Context, id int) (*Client, error)
}

type RepositoryPostgres struct {
	DB *sql.DB
}

func NewRepositoryPostgres(db *sql.DB) *RepositoryPostgres {
	return &RepositoryPostgres{DB: db}
}

func (r *RepositoryPostgres) FindOneByID(ctx context.Context, id int) (*Client, error) {
	var client Client

	row := r.DB.QueryRowContext(ctx, "SELECT * FROM clients WHERE id = $1", id)

	err := row.Scan(&client.ID, &client.AccountLimit, &client.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrClientNotFound
		}
		return nil, err
	}

	return &client, nil
}
