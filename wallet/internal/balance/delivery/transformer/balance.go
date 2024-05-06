package transformer

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/contract"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/entity"
	"net/http"
)

const (
	InvalidUserIDErrMsg = "userID is invalid"
)

type builder struct {
	filters entity.Filters
	req     *http.Request
	err     error
}

func (b *builder) SetUserID() contract.Transformer {
	const op = "balanceTransformer.SetUserID"
	if b.err != nil {
		return b
	}

	vars := mux.Vars(b.req)
	val, ok := vars["userID"]
	if !ok {
		b.err = richerror.New(op).WithErr(errors.New(InvalidUserIDErrMsg))
		return b
	}
	b.filters.UserID = val
	return b
}

func (b *builder) Transform() (entity.Filters, error) {
	return b.filters, b.err
}

func NewBalanceTransformer(r *http.Request) contract.Transformer {
	return &builder{req: r}
}
