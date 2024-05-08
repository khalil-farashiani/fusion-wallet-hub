package main

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/adapter/rabbitmq"
	redis "github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/adapter/redis"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/config"
	httpServer "github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/delivery/http"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/repository/migrator"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/repository/mysql"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/repository/mysql/mysql_redeem"
	redisRepo "github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/repository/redis"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/service"
	"log"
	"net/http"
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
		// curl http://localhost:8090/debug/pprof/goroutine --output goroutine.o
		//  go tool pprof -http=:8086 ./goroutine.o
		http.ListenAndServe(":8090", nil)
	}()

	args := os.Args
	if len(args) < 2 {
		panic(notEnoughArgumentErrMsg)
	}
	cfg := config.Load(args[ConfigAddressArg])

	mgr := migrator.New(cfg.Mysql)
	mgr.Up()

	discountSvc := setupServices(cfg)
	server := httpServer.New(cfg, discountSvc)

	go func() {
		server.Serve()
	}()

	sig := waitExitSignal()
	log.Println(sig.String())
}

func setupServices(cfg config.Config) service.Service {
	MysqlRepo := mysql.New(cfg.Mysql)
	redisAdapter := redis.New(cfg.Cache)

	bSQL := mysql_redeem.New(MysqlRepo)
	broker := rabbitmq.New(cfg.Broker)
	cache := redisRepo.New(redisAdapter)

	rs := service.New(bSQL, cache, broker)
	return rs
}

// waitExitSignal get os signals
func waitExitSignal() os.Signal {
	quit := make(chan os.Signal, 6)
	signal.Notify(quit, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	return <-quit
}
