package dto

import "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/entity"

type UserBalanceResponse struct {
	Amount int64 `json:"amount"`
}

func ToUserBalanceResponse(in entity.Balance) UserBalanceResponse {
	return UserBalanceResponse{
		Amount: in.Amount,
	}
}
