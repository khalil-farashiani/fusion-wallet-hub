package entity

type Event string

const (
	AddAmountEvent Event = "user_discount"
)

type BrokerAddAmountEventData struct {
	UserID string `json:"user_id"`
	Amount uint64 `json:"amount"`
}
