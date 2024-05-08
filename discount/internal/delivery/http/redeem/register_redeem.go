package handler

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/delivery/dto"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/delivery/transformer"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/http_helper"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/http_msg"
	"net/http"
)

func (h *Handler) RegisterRedeem(w http.ResponseWriter, r *http.Request) {
	params, err := transformer.NewDiscountTransformer(r).
		SetUserID().
		SetTitle().
		Transform()
	if err != nil {
		msg, code := http_msg.Error(err)
		http_helper.JSONErr(w, code, msg)
		return
	}

	if err = h.discountSvc.SetUserRedeem(params.UserID,
		dto.FromFilterEntityToRedeemReportEntity(params)); err != nil {
		msg, code := http_msg.Error(err)
		http_helper.JSONErr(w, code, msg)
		return
	}

	http_helper.JSON(w, http.StatusOK, map[string]string{
		"message": "success",
	})
}
