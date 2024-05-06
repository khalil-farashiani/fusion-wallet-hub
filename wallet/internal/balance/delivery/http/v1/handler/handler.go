package handler

import "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/service"

type Handler struct {
	balanceSvc service.Service
}

func New(balanceSvc service.Service) Handler {
	return Handler{
		balanceSvc: balanceSvc,
	}
}
