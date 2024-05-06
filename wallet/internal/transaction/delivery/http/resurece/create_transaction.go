package resurece

import "time"

type CreateTransactionRequest struct {
	AccountID string `json:"account_id"`
	Type      uint64 `json:"type"`
	Amount    uint64 `json:"amount"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
