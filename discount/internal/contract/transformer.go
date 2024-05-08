package contract

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/delivery/resource"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/entity"
)

type Transformer interface {
	SetStatus() Transformer
	SetTitle() Transformer
	SetRedeemBody() Transformer
	SetUserID() Transformer
	GetRedeemBody() (resource.CreateRedeemBody, error)
	Transform() (entity.Filter, error)
}
