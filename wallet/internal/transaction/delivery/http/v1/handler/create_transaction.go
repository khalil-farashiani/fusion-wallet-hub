package handler

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/http_helper"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/http_msg"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/delivery/transformer"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/dto"
	"net/http"
)

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	params, err := transformer.NewTransactionTransformer(r).
		SetTransactionCreateBody().
		GetTransactionCreateBody()

	err = h.transactionSvc.CreateSingle(dto.CreateTransactionBodyToTransactionEntity(params))
	if err != nil {
		msg, code := http_msg.Error(err)
		http_helper.JSONErr(w, code, msg)
		return
	}

	http_helper.JSON(w, http.StatusCreated, map[string]string{
		"message": "success",
	})
}
