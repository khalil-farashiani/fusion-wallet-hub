package contract

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/entity"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/delivery/http/resurece"
)

type Transformer interface {
	SetUserID() Transformer
	SetDateTimeFilter() Transformer
	SetCursor() Transformer
	SetLimit() Transformer
	SetTransactionCreateBody() Transformer
	GetTransactionCreateBody() (resurece.CreateTransactionRequest, error)
	Transform() (entity.Filters, error)
}
