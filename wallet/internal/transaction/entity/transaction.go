package entity

import "time"

type TransactionType string

type Transaction struct {
	ID        uint            `json:"id"`
	AccountID string          `json:"account_id"`
	Type      TransactionType `json:"type"`
	Amount    uint64          `json:"amount"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
