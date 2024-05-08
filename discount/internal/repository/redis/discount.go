package redis

import (
	"context"
	"fmt"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
)

// TODO: we can use env for redis prefixes
const discountPrefix = "discount"

func (d DB) SetInitialRedeemCount(key string, initialValue uint64) error {
	const op = richerror.Op("cache.SetInitialRedeemCount")

	_, err := d.adapter.Client().
		Set(context.Background(),
			fmt.Sprintf("%s%s", discountPrefix, key),
			initialValue,
			0).
		Result()
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}

	return nil
}

func (d DB) DecrementRedeemCount(key string) (int64, error) {
	const op = richerror.Op("cache.DecrementRedeemCount")

	newValue, err := d.adapter.Client().
		Decr(context.Background(),
			fmt.Sprintf("%s%s", discountPrefix, key)).
		Result()
	if err != nil {
		return 0, richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}

	return newValue, nil
}
