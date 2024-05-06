package handler

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/http_helper"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/http_msg"
	transformer2 "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/delivery/transformer"
	"net/http"
)

func (h *Handler) GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	params, err := transformer2.NewTransactionTransformer(r).
		SetCursor().
		SetDateTimeFilter().
		SetLimit().
		Transform()

	result, err := h.transactionSvc.GetAll(*params.Paginate)
	if err != nil {
		msg, code := http_msg.Error(err)
		http_helper.JSONErr(w, code, msg)
		return
	}

	http_helper.JSON(w, http.StatusCreated, result)
}
