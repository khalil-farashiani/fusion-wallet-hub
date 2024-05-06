package dto

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/delivery/http/resurece"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/entity"
	"time"
)

func CreateTransactionBodyToTransactionEntity(request resurece.CreateTransactionRequest) entity.Transaction {
	return entity.Transaction{
		AccountID: request.AccountID,
		Type:      entity.TransactionType(request.Type),
		Amount:    request.Amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
