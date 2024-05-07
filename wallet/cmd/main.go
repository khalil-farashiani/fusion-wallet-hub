package main

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/adapter/rabbitmq"
	balanceService "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/service"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/config"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/delivery/consumer"
	httpserver "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/delivery/http"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/repository/migrator"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/repository/mysql"
	balanceMysql "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/repository/mysql/mysql_balance"
	transactionMysql "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/repository/mysql/mysql_transaction"
	transactionService "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/service"
	"net/http"

	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	ConfigAddressArg        = 1
	notEnoughArgumentErrMsg = "error: config address is a required argument"
)

func main() {
	go func() {
		// curl http://localhost:8099/debug/pprof/goroutine --output goroutine.o
		//  go tool pprof -http=:8086 ./goroutine.o
		http.ListenAndServe(":8099", nil)
	}()

	args := os.Args
	if len(args) < 2 {
		panic(notEnoughArgumentErrMsg)
	}
	cfg := config.Load(args[ConfigAddressArg])

	mgr := migrator.New(cfg.Mysql)
	mgr.Up()

	balanceSvc, transactionSvc, broker := setupServices(cfg)
	server := httpserver.New(cfg, balanceSvc, transactionSvc)
	consumer := consumer.New(cfg, balanceSvc, transactionSvc, broker)

	go func() {
		server.Serve()
	}()

	go func() {
		consumer.Start()
	}()

	sig := waitExitSignal()
	log.Println(sig.String())
}

func setupServices(cfg config.Config) (balanceService.Service, transactionService.Service, rabbitmq.Adapter) {
	MysqlRepo := mysql.New(cfg.Mysql)

	bSQL := balanceMysql.New(MysqlRepo)
	tSQL := transactionMysql.New(MysqlRepo)

	bs := balanceService.New(bSQL)
	ts := transactionService.New(tSQL)

	broker := rabbitmq.New(cfg.Broker)
	return bs, ts, broker
}

// waitExitSignal get os signals
func waitExitSignal() os.Signal {
	quit := make(chan os.Signal, 6)
	signal.Notify(quit, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	return <-quit
}
