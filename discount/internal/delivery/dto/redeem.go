package dto

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/delivery/resource"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/entity"
	"time"
)

func FromCreateRedeemToRedeemEntity(req resource.CreateRedeemBody) entity.Redeem {
	return entity.Redeem{
		Title:     req.Title,
		Amount:    req.Amount,
		Quantity:  req.Quantity,
		CreatedAt: time.Now(),
	}
}

func FromFilterEntityToRedeemReportEntity(from entity.Filter) entity.RedeemReport {
	return entity.RedeemReport{
		Title:  from.Title,
		Status: from.Status,
	}
}
