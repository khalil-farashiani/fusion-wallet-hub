package http

import (
	"fmt"
	"github.com/gorilla/mux"
	balanceHdr "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/delivery/http/v1/handler"
	BalanceService "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/balance/service"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/config"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/logger"
	transactionHdr "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/delivery/http/v1/handler"
	transactionService "github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/transaction/service"
	"log"
	"net/http"
	"time"
)

type Server struct {
	config config.Config

	balanceHandler     balanceHdr.Handler
	transactionHandler transactionHdr.Handler

	Router *mux.Router
}

func New(cfg config.Config, balanceSvc BalanceService.Service, transactionSvc transactionService.Service) Server {
	return Server{
		config: cfg,

		Router: mux.NewRouter(),

		balanceHandler:     balanceHdr.New(balanceSvc),
		transactionHandler: transactionHdr.New(transactionSvc),
	}
}

func (s Server) Serve() {
	lg := logger.Logger.Named("http-server")
	s.Router.Use(loggingMiddleware(lg))

	s.Router.HandleFunc("/health-check", healthCheckHandler)
	s.balanceHandler.SetBalanceRoute(s.Router)
	s.transactionHandler.SetTransactionRoute(s.Router)
	address := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	fmt.Printf("start wallet server on %s\n", address)

	srv := &http.Server{
		Handler: s.Router,
		Addr:    fmt.Sprintf("127.0.0.1%s", address),
		//enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
