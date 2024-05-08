package contract

import "github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/entity"

type Publisher interface {
	Publish(event entity.Event, payload string) error
}
