package entity

import "time"

type TransactionType int

const (
	Debit  TransactionType = iota // 0
	Credit                        // 1
)

// String returns the string representation of the TransactionType.
func (t TransactionType) String() string {
	types := []string{
		"DEBIT",
		"CREDIT",
	}
	if int(t) < len(types) {
		return types[t]
	}

	//TODO: handle invalid types gracefully
	return "UNKNOWN"
}

type Transaction struct {
	ID        uint            `json:"id"`
	AccountID string          `json:"account_id"`
	Type      TransactionType `json:"type"`
	Amount    uint64          `json:"amount"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
