package dto

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/entity"
	"time"
)

func CreateTransactionEntity(userID string, amount uint64) entity.Transaction {
	return entity.Transaction{
		AccountID: userID,
		Type:      "IN",
		Amount:    amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
