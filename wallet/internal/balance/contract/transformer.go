package contract

import "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/entity"

type Transformer interface {
	SetUserID() Transformer
	Transform() (entity.Filters, error)
}
