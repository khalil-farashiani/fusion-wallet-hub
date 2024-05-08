package transformer

import (
	"encoding/json"
	"errors"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/contract"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/delivery/resource"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/entity"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"net/http"
	"strings"
)

const (
	InvalidStatusErrMsg    = "status is invalid"
	invalidJSONBodyErrMsg  = "invalid JSON body"
	RequiredUserIDErrMsg   = "user_id is required field"
	RequiredStatusIDErrMsg = "status is required field"
	RequiredTitleIDErrMsg  = "title is required field"
)

type builder struct {
	filters      entity.Filter
	createRedeem resource.CreateRedeemBody
	req          *http.Request
	err          error
}

func (b *builder) validateStatus(status string) (string, error) {
	var validStatuses = map[string]string{
		"invalid": "USED",
		"valid":   "NEW",
	}
	val, ok := validStatuses[status]
	if !ok {
		return "", errors.New(InvalidStatusErrMsg)
	}
	return val, nil
}

func (b *builder) GetRedeemBody() (resource.CreateRedeemBody, error) {
	return b.createRedeem, b.err
}

func (b *builder) SetStatus() contract.Transformer {
	const op = "discountTransformer.SetStatus"
	if b.err != nil {
		return b
	}

	status := b.req.URL.Query().Get("status")
	if status != "" {
		mappedStatus, err := b.validateStatus(status)
		if err != nil {
			b.err = richerror.New(op).WithErr(err)
			return b
		}
		b.filters.Status = mappedStatus
	} else {
		b.err = richerror.New(op).WithErr(errors.New(RequiredStatusIDErrMsg))
	}
	return b
}

func (b *builder) SetTitle() contract.Transformer {
	const op = "discountTransformer.SetStatus"
	if b.err != nil {
		return b
	}

	title := b.req.URL.Query().Get("title")
	if title != "" {
		b.filters.Title = title
	} else {
		b.err = richerror.New(op).WithErr(errors.New(RequiredTitleIDErrMsg))
	}
	return b
}

func (b *builder) SetUserID() contract.Transformer {
	const op = "discountTransformer.SetUserID"
	if b.err != nil {
		return b
	}

	userID := b.req.URL.Query().Get("user_id")
	if userID != "" {
		b.filters.UserID = userID
	} else {
		b.err = richerror.New(op).WithErr(errors.New(RequiredUserIDErrMsg))
	}
	return b
}

func (b *builder) validateTransactionCreateBody(req resource.CreateRedeemBody) error {
	if req.Amount == 0 || strings.TrimSpace(req.Title) == "" || req.Quantity == 0 {
		return errors.New(invalidJSONBodyErrMsg)
	}
	return nil
}

func (b *builder) SetRedeemBody() contract.Transformer {
	const op = "discountTransformer.SetRedeemBody"
	if b.err != nil {
		return b
	}

	err := json.NewDecoder(b.req.Body).Decode(&b.createRedeem)
	if err != nil {
		b.err = richerror.New(op).WithErr(errors.New(invalidJSONBodyErrMsg))
		return b
	}
	b.err = b.validateTransactionCreateBody(b.createRedeem)
	return b
}

func (b *builder) Transform() (entity.Filter, error) {
	return b.filters, b.err
}

func NewDiscountTransformer(r *http.Request) contract.Transformer {
	return &builder{req: r}
}
