package handler

import "github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/service"

type Handler struct {
	discountSvc service.Service
}

func New(discountSvc service.Service) Handler {
	return Handler{
		discountSvc: discountSvc,
	}
}
