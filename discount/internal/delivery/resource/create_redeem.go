package resource

type CreateRedeemBody struct {
	Title    string `json:"title"`
	Amount   uint64 `json:"amount"`
	Quantity uint64 `json:"quantity"`
}
