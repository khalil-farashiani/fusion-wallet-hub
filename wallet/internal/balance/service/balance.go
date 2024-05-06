package service

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/entity"
)

type Repository interface {
	GetUserBalance(userID string) (entity.Balance, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) GetUserBalance(userID string) (entity.Balance, error) {
	const op = "service.GetUserBalance"

	balance, err := s.repo.GetUserBalance(userID)
	if err != nil {
		//might better sanitize request
		return entity.Balance{},
			richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}

	return balance, err
}
