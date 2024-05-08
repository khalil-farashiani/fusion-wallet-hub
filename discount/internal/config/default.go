package config

import "time"

var defaultConfig = map[string]interface{}{
	"application.graceful_shutdown_timeout": time.Second * 10,
}
