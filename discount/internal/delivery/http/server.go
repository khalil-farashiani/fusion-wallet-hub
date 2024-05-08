package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/config"
	handler "github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/delivery/http/redeem"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/logger"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/service"
	"log"
	"net/http"
	"time"
)

type Server struct {
	config config.Config

	discountHandler handler.Handler

	Router *mux.Router
}

func New(cfg config.Config, redeemSvc service.Service) Server {
	return Server{
		config: cfg,

		Router: mux.NewRouter(),

		discountHandler: handler.New(redeemSvc),
	}
}

func (s Server) Serve() {
	lg := logger.Logger.Named("http-server")
	s.Router.Use(loggingMiddleware(lg))

	s.Router.HandleFunc("/health-check", healthCheckHandler)

	s.discountHandler.SetDiscountRoute(s.Router)
	address := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	fmt.Printf("start discount server on %s\n", address)

	srv := &http.Server{
		Handler: s.Router,
		Addr:    fmt.Sprintf("127.0.0.1%s", address),
		//enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
