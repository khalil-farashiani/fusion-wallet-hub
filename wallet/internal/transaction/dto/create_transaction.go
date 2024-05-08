package dto

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/delivery/http/resurece"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/entity"
	"time"
)

func typeIntToValue(typee uint64) entity.TransactionType {
	var mapTypeIntToString = map[uint64]string{
		1: "IN",
		2: "OUT",
	}
	return entity.TransactionType(mapTypeIntToString[typee])
}

func CreateTransactionBodyToTransactionEntity(request resurece.CreateTransactionRequest) entity.Transaction {
	return entity.Transaction{
		AccountID: request.AccountID,
		Type:      typeIntToValue(request.Type),
		Amount:    request.Amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
