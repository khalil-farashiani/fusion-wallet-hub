package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"strings"
)

func Load(configPath string) Config {
	// Global koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
	var k = koanf.New(".")

	k.Load(confmap.Provider(defaultConfig, "."), nil)

	// Load YAML config and merge into the previously loaded config.
	k.Load(file.Provider(configPath), yaml.Parser())

	k.Load(env.Provider("DISCOUNT_", ".", func(s string) string {
		str := strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "DISCOUNT_")), "_", ".", -1)

		return strings.Replace(str, "..", "_", -1)
	}), nil)

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}

	return cfg
}
