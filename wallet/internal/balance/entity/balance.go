package entity

import "time"

type Balance struct {
	ID        uint   `json:"id"`
	AccountID string `json:"account_id"`
	Amount    int64  `json:"amount"`

	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
