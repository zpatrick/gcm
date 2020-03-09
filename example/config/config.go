package config

import (
	"flag"

	"github.com/zpatrick/gcm"
)

// ConfigFilePath denotes the location of the configuration file.
const ConfigFilePath = "config.yaml"

// Collection of canonical configuration keys.
const (
	KeyRedisHost = "redis.host"
	KeyRedisPort = "redis.port"
)

// Mux returns the configuration mux for the application.
func Mux() (*gcm.Mux, error) {
	flagProvider := gcm.NewFlagProvider(flag.NewFlagSet("app", flag.ContinueOnError))
	fileProvider := gcm.NewFileProvider(ConfigFilePath, gcm.YAMLParser, gcm.ReloadNever)
	envProvider := gcm.NewEnvironmentProvider()

	m := &gcm.Mux{
		Providers: map[string]gcm.Provider{
			KeyRedisHost: &gcm.StringProviderSchema{
				Default:       "localhost",
				DefaultIsZero: true,
				Provider: gcm.MultiStringProvider{
					flagProvider.String("redis-host", "localhost", "redis host", false),
					fileProvider.String("redis", "host"),
					envProvider.String("APP_REDIS_HOST"),
				},
			},
			KeyRedisPort: &gcm.IntProviderSchema{
				Default:  6379,
				Validate: gcm.ValidateIntBetween(0, 65535),
				Provider: gcm.MultiIntProvider{
					flagProvider.Int("redis-port", 6379, "redis port", false),
					fileProvider.Int("redis", "port"),
					envProvider.Int("APP_REDIS_PORT"),
				},
			},
		},
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}
