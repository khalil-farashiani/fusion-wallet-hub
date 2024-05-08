package service

import (
	"encoding/json"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/contract"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/entity"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
)

type Repository interface {
	CreateUserRedeemRecord(redeem entity.RedeemReport) error
	IsUserUseRedeemBefore(redeem entity.RedeemReport) (bool, error)
	CreateRedeem(redeem entity.Redeem) error
	GetReports(status string) ([]entity.RedeemReport, error)
	GetRedeem(title string) (entity.Redeem, error)
}

type DiscountCache interface {
	SetInitialRedeemCount(key string, initialValue uint64) error
	DecrementRedeemCount(key string) (int64, error)
}

type Service struct {
	repo   Repository
	broker contract.Publisher
	cache  DiscountCache
}

func New(repo Repository, cache DiscountCache, broker contract.Publisher) Service {
	return Service{repo: repo, broker: broker, cache: cache}
}

func (s Service) toStatusMapper(isUserUseRedeemBefore bool) string {
	if isUserUseRedeemBefore {
		return "USED"
	}
	return "NEW"
}

func (s Service) GetRedeemReport(status string) ([]entity.RedeemReport, error) {
	const op = "service.GetRedeemReport"

	reports, err := s.repo.GetReports(status)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}

	return reports, nil
}

// we can also use protobuf
func (s Service) prepareAddBalancePayload(userID string, Amount uint64) string {
	data := entity.BrokerAddAmountEventData{
		UserID: userID,
		Amount: Amount,
	}

	bData, _ := json.Marshal(data)
	return string(bData)
}

func (s Service) SetUserRedeem(userID string, redeemRep entity.RedeemReport) error {
	const op = "service.SetUserRedeem"

	//check amount from cache
	redeemRep.UserId = userID
	redeemRep.Status = "NEW"
	isUserUseRedeemBefore, err := s.repo.IsUserUseRedeemBefore(redeemRep)
	if err != nil {
		//might better sanitize request
		return richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}
	if !isUserUseRedeemBefore {
		redeem, err := s.repo.GetRedeem(redeemRep.Title)
		if err != nil {
			return err
		}
		if err := s.broker.Publish(entity.AddAmountEvent, s.prepareAddBalancePayload(userID, redeem.Amount)); err != nil {
			return richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
		}
		redeemRep.Amount = redeem.Amount
	}
	redeemRep.Status = s.toStatusMapper(isUserUseRedeemBefore)

	return s.repo.CreateUserRedeemRecord(redeemRep)
}

func (s Service) NewRedeem(redeem entity.Redeem) error {
	const op = "service.SetUserRedeem"

	err := s.repo.CreateRedeem(redeem)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}
	if err = s.cache.SetInitialRedeemCount(redeem.Title, redeem.Amount); err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.Unexpected)
	}
	return nil
}
