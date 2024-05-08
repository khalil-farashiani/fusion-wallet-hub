package entity

import "time"

type RedeemReport struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
	Amount uint64 `json:"amount"`
	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
}

type Redeem struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Amount   uint64 `json:"amount"`
	Quantity uint64 `json:"quantity"`

	CreatedAt time.Time `json:"created_at"`
}
