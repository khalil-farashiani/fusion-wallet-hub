package handler

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/http_helper"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/http_msg"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/delivery/dto"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/delivery/transformer"
	"net/http"
)

func (h *Handler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	params, err := transformer.NewBalanceTransformer(r).
		SetUserID().
		Transform()
	if err != nil {
		msg, code := http_msg.Error(err)
		http_helper.JSONErr(w, code, msg)
		return
	}

	result, err := h.balanceSvc.GetUserBalance(params.UserID)
	if err != nil {
		msg, code := http_msg.Error(err)
		http_helper.JSONErr(w, code, msg)
		return
	}

	http_helper.JSON(w, http.StatusOK, dto.ToUserBalanceResponse(result))
}
