package config

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/adapter/rabbitmq"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/adapter/redis"
	"github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/repository/mysql"
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
	Cache       redis.Config    `koanf:"redis"`
	Mysql       mysql.Config    `koanf:"mysql"`
}
