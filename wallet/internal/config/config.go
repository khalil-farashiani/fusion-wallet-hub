package config

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/adapter/rabbitmq"
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/repository/mysql"
	"time"
)

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct {
	Application Application     `koanf:"application"`
	HTTPServer  HTTPServer      `koanf:"http_server"`
	Broker      rabbitmq.Config `koanf:"broker"`
	Mysql       mysql.Config    `koanf:"mysql"`
}
