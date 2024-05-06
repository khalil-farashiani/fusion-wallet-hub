package service

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	generalEntity "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/entity"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/entity"
	"time"
)

type Repository interface {
	CreateTransaction(transaction entity.Transaction) error
	GetUserTransactions(userID string, pg generalEntity.Paginate) ([]entity.Transaction, error)
	GetTransactions(pg generalEntity.Paginate) ([]entity.Transaction, error)
}

type Service struct {
	repo Repository
}

func (s *Service) UserTransactions(userID string, pg generalEntity.Paginate) ([]entity.Transaction, error) {
	const op = "service.UserTransactions"

	result, err := s.repo.GetUserTransactions(userID, pg)
	if err != nil {
		//might better sanitize request
		return nil,
			richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}
	return result, nil
}

func (s *Service) GetAll(pg generalEntity.Paginate) ([]entity.Transaction, error) {
	const op = "service.GetTransactions"

	result, err := s.repo.GetTransactions(pg)
	if err != nil {
		return nil,
			richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}
	return result, nil
}

func (s *Service) CreateSingle(tx entity.Transaction) error {
	const op = "service.CreateSingle"
	//better we move time functionality to infrastructure layer
	tx.CreatedAt = time.Now()
	tx.UpdatedAt = time.Now()

	if err := s.repo.CreateTransaction(tx); err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}
	return nil
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
