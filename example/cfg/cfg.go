package cfg

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/zpatrick/gcm"
	"gopkg.in/yaml.v2"
)

// ConfigFilePath denotes the location of the configuration file.
const ConfigFilePath = "config.yaml"

// Collection of canonical configuration keys.
const (
	KeyRedisHost = "redis.host"
	KeyRedisPort = "redis.port"
)

type YAMLFileConfig struct {
	Redis struct {
		Host *string `json:"host" yaml:"host"`
		Port *int    `json:"port" yaml:"port"`
	} `json:"redis" yaml:"redis"`
}

func loadYAMLFile(path string) (YAMLFileConfig, error) {
	content, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		return YAMLFileConfig{}, err
	}

	var yamlFile YAMLFileConfig
	if err := yaml.Unmarshal(content, &yamlFile); err != nil {
		return YAMLFileConfig{}, err
	}

	return yamlFile, nil
}

// Load returns the configuration mux for the application.
func Load() (*gcm.Mux, error) {
	yamlFile, err := loadYAMLFile(ConfigFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load '%s': %w", ConfigFilePath, err)
	}

	flagProvider := gcm.NewFlagProvider(flag.NewFlagSet("app", flag.ContinueOnError))
	envProvider := gcm.NewEnvironmentProvider()

	m := &gcm.Mux{
		Providers: map[string]gcm.Provider{
			KeyRedisHost: &gcm.StringProviderSchema{
				Default: "localhost",
				Provider: gcm.MultiStringProvider{
					flagProvider.String("redis-host", "localhost", "redis host", false),
					envProvider.String("APP_REDIS_HOST"),
					gcm.OptionalStaticString(yamlFile.Redis.Host),
				},
			},
			KeyRedisPort: &gcm.IntProviderSchema{
				Default:  6379,
				Validate: gcm.ValidateIntBetween(0, 65535),
				Provider: gcm.MultiIntProvider{
					flagProvider.Int("redis-port", 6379, "redis port", false),
					envProvider.Int("APP_REDIS_PORT"),
					gcm.OptionalStaticInt(yamlFile.Redis.Port),
				},
			},
		},
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}
