package clients

import (
	"context"
	"time"
)

type Service struct {
	Repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{Repository: repository}
}

type ClientExtract struct {
	Date         string `json:"date"`
	AccountLimit int    `json:"account_limit"`
	Balance      int    `json:"balance"`
}

func (s *Service) GetClientExtract(ctx context.Context, id int) (*ClientExtract, error) {
	client, err := s.Repository.FindOneByID(ctx, id, nil)
	if err != nil {
		return nil, err
	}

	return &ClientExtract{
		Date:         time.Now().Format(time.RFC3339),
		AccountLimit: client.AccountLimit,
		Balance:      client.Balance,
	}, nil
}
