package handler

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/service"
)

type Handler struct {
	transactionSvc service.Service
}

func New(transactionSvc service.Service) Handler {
	return Handler{
		transactionSvc: transactionSvc,
	}
}
