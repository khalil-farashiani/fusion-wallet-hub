module github.com/khalil-farashiani/fusion-wallet-hub/discount

go 1.21.6

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/khalil-farashiani/fusion-wallet-hub/pkg v0.0.0
	github.com/knadh/koanf/parsers/yaml v0.1.0
	github.com/knadh/koanf/providers/confmap v0.1.0
	github.com/knadh/koanf/providers/env v0.1.0
	github.com/knadh/koanf/providers/file v0.1.0
	github.com/knadh/koanf/v2 v2.1.1
	github.com/rabbitmq/amqp091-go v1.9.0
	github.com/redis/go-redis/v9 v9.5.1
	github.com/rubenv/sql-migrate v1.6.1
	github.com/stretchr/testify v1.9.0
	go.uber.org/zap v1.27.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-gorp/gorp/v3 v3.1.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/khalil-farashiani/fusion-wallet-hub/pkg v0.0.0 => ../pkg
