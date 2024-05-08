package consumer

import (
	"encoding/json"
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/adapter/rabbitmq"
	BalanceService "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/service"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/config"
	dto "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/delivery/dto"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/entity"
	transactionService "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/service"
	"log"
)

type Consumer struct {
	config         config.Config
	balanceSvc     BalanceService.Service
	transactionSvc transactionService.Service
	broker         rabbitmq.Adapter
}

func New(cfg config.Config, balanceSvc BalanceService.Service, transactionSvc transactionService.Service, broker rabbitmq.Adapter) Consumer {
	return Consumer{
		config:         cfg,
		balanceSvc:     balanceSvc,
		transactionSvc: transactionSvc,
		broker:         broker,
	}
}

func (s Consumer) Start() {
	const op = "consumer.Start"
	msgs, err := s.broker.Consume(entity.AddAmountEvent)
	if err != nil {
		log.Fatal(richerror.New(op).WithMessage("Failed to register a consumer").WithErr(err))
	}

	for d := range msgs {
		var event entity.BrokerAddAmountEventData
		if err := json.Unmarshal(d.Body, &event); err != nil {
			//TODO: log fails
			d.Nack(false, true) // requeue the message
			continue
		}

		if err := s.balanceSvc.IncreaseUserBalance(event.UserID, event.Amount); err != nil {
			d.Nack(false, true) // requeue the message
			continue
		}
		if err := s.transactionSvc.CreateSingle(dto.CreateTransactionEntity(event.UserID, uint64(event.Amount))); err != nil {
			d.Nack(false, true) // requeue the message
			continue
		}
		d.Ack(false)
	}
}
