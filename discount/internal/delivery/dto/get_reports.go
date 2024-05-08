package dto

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/entity"
	"time"
)

type GetReportResponse struct {
	ID     uint   `json:"id"`
	UserId string `json:"user_id"`
	Amount uint64 `json:"amount"`

	CreatedAt time.Time `json:"created_at"`
}

func ToGetReportResponse(report entity.RedeemReport) GetReportResponse {
	return GetReportResponse{
		ID:        report.ID,
		UserId:    report.UserId,
		Amount:    report.Amount,
		CreatedAt: report.CreatedAt,
	}
}

func ToGetReportsResponse(reports []entity.RedeemReport) []GetReportResponse {
	var result = make([]GetReportResponse, 0, len(reports))
	for _, r := range reports {
		res := ToGetReportResponse(r)
		result = append(result, res)
	}
	return result
}
