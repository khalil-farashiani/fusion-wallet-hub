package service

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/entity"
	"sync"
)

type Repository interface {
	GetUserBalance(userID string) (entity.Balance, error)
	IncreaseBalance(userID string, amount uint64) error
}

type Service struct {
	repo  Repository
	mutex *sync.Mutex
}

func New(repo Repository) Service {
	return Service{repo: repo, mutex: &sync.Mutex{}}
}

func (s Service) GetUserBalance(userID string) (entity.Balance, error) {
	return s.repo.GetUserBalance(userID)
}

func (s Service) IncreaseUserBalance(userID string, amount int64) error {
	const op = "service.IncreaseUserBalance"
	s.mutex.Lock()
	defer s.mutex.Unlock() // Ensure the lock is released after the function execution

	balance, err := s.repo.GetUserBalance(userID)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}

	balance.Amount += amount
	err = s.repo.IncreaseBalance(userID, uint64(balance.Amount))
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}

	return nil
}
