package transformer

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/entity"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/contract"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/delivery/http/resurece"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	invalidUserIDErrMsg         = "userID is invalid"
	invalidToDateFormatErrMsg   = "invalid 'to' date format"
	invalidFromDateFormatErrMsg = "invalid 'from' date format"
	invalidCursorErrMsg         = "invalid cursor format"
	invalidLimitErrMsg          = "invalid limit format"
	invalidJSONBodyErrMsg       = "invalid json body"

	defaultLimitValue = 100
)

type builder struct {
	filters               entity.Filters
	createTransactionBody resurece.CreateTransactionRequest
	req                   *http.Request
	err                   error
}

func (b *builder) SetUserID() contract.Transformer {
	const op = "balanceTransformer.SetUserID"
	if b.err != nil {
		return b
	}

	vars := mux.Vars(b.req)
	val, ok := vars["userID"]
	if !ok {
		b.err = richerror.New(op).WithErr(errors.New(invalidUserIDErrMsg))
		return b
	}
	b.filters.UserID = val
	return b
}

func (b *builder) SetDateTimeFilter() contract.Transformer {
	const op = "transactionTransformer.SetDateTimeFilter"
	if b.err != nil {
		return b
	}

	// Parse 'from' date
	from := b.req.URL.Query().Get("from")
	if from != "" {
		fromDate, err := time.Parse("2006-01-02", from)
		if err != nil {
			b.err = richerror.New(op).WithErr(errors.New(invalidFromDateFormatErrMsg))
			return b
		}
		b.filters.Paginate.AfterDateTime = fromDate
	}

	// Parse 'to' date
	to := b.req.URL.Query().Get("to")
	if to != "" {
		toDate, err := time.Parse("2006-01-02", to)
		if err != nil {
			b.err = richerror.New(op).WithErr(errors.New(invalidToDateFormatErrMsg))
			return b
		}
		b.filters.Paginate.BeforeDateTime = toDate
	}

	return b
}

func (b *builder) SetCursor() contract.Transformer {
	const op = "transactionTransformer.SetCursor"
	if b.err != nil {
		return b
	}

	cursor := b.req.URL.Query().Get("cursor")
	if cursor != "" {
		cursorInt, err := strconv.Atoi(cursor)
		if err != nil {
			b.err = richerror.New(op).WithErr(errors.New(invalidCursorErrMsg))
			return b
		}
		b.filters.Paginate.Cursor = int64(cursorInt)
	}
	return b
}

func (b *builder) SetLimit() contract.Transformer {
	const op = "transactionTransformer.SetLimit"
	if b.err != nil {
		return b
	}

	limit := b.req.URL.Query().Get("limit")
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			b.err = richerror.New(op).WithErr(errors.New(invalidLimitErrMsg))
			return b
		}
		b.filters.Paginate.Limit = int64(limitInt)
		return b
	}
	b.filters.Paginate.Limit = defaultLimitValue
	return b
}

func validateTransactionCreateBody(req resurece.CreateTransactionRequest) error {
	const op = "transactionTransformer.validateTransactionCreateBody"

	if (req.Amount == 0) || (req.Type != 1 && req.Type != 2) || strings.TrimSpace(req.AccountID) == "" {
		return richerror.New(op).WithErr(errors.New(invalidJSONBodyErrMsg))
	}
	return nil
}

func (b *builder) SetTransactionCreateBody() contract.Transformer {
	const op = "transactionTransformer.SetTransactionCreateBody"
	if b.err != nil {
		return b
	}

	err := json.NewDecoder(b.req.Body).Decode(&b.createTransactionBody)
	if err != nil {
		b.err = richerror.New(op).WithErr(errors.New(invalidJSONBodyErrMsg))
		return b
	}
	b.err = validateTransactionCreateBody(b.createTransactionBody)

	return b
}

func (b *builder) GetTransactionCreateBody() (resurece.CreateTransactionRequest, error) {
	return b.createTransactionBody, b.err
}

func NewTransactionTransformer(r *http.Request) contract.Transformer {
	return &builder{req: r}
}

func (b *builder) Transform() (entity.Filters, error) {
	return b.filters, b.err
}
