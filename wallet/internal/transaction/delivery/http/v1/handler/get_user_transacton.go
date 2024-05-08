package handler

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/http_helper"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/http_msg"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/delivery/transformer"
	"net/http"
)

func (h *Handler) GetUserTransaction(w http.ResponseWriter, r *http.Request) {
	params, err := transformer.NewTransactionTransformer(r).
		SetUserID().
		SetCursor().
		SetDateTimeFilter().
		SetLimit().
		Transform()

	result, err := h.transactionSvc.UserTransactions(params.UserID, params.Paginate)
	if err != nil {
		msg, code := http_msg.Error(err)
		http_helper.JSONErr(w, code, msg)
		return
	}

	http_helper.JSON(w, http.StatusCreated, result)
}
