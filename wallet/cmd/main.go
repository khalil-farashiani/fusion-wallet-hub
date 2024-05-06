package main

import (
	balanceService "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/service"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/config"
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
	ConfigAddressArg        = 2
	notEnoughArgumentErrMsg = "error: config address is a required argument"
)

func main() {
	go func() {
		// curl http://localhost:8099/debug/pprof/goroutine --output goroutine.o
		//  go tool pprof -http=:8086 ./goroutine.o
		http.ListenAndServe(":8099", nil)
	}()

	args := os.Args
	if len(args) < 3 {
		panic(notEnoughArgumentErrMsg)
	}
	cfg := config.Load(args[ConfigAddressArg])

	mgr := migrator.New(cfg.Mysql)
	mgr.Up()

	authSvc, userSvc := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc)

	go func() {
		server.Serve()
	}()

	sig := waitExitSignal()
	log.Println(sig.String())
}

func setupServices(cfg config.Config) (balanceService.Service, transactionService.Service) {
	MysqlRepo := mysql.New(cfg.Mysql)

	bSQL := balanceMysql.New(MysqlRepo)
	tSQL := transactionMysql.New(MysqlRepo)

	bs := balanceService.New(bSQL)
	ts := transactionService.New(tSQL)
	return bs, ts
}

// waitExitSignal get os signals
func waitExitSignal() os.Signal {
	quit := make(chan os.Signal, 6)
	signal.Notify(quit, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	return <-quit
}
