package service

import (
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
