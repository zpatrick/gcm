package gcm

import (
	"fmt"
	"os"
	"strconv"
)

type EnvironmentProvider struct{}

func NewEnvironmentProvider() *EnvironmentProvider {
	return &EnvironmentProvider{}
}

func (e *EnvironmentProvider) get(envvar string) (string, error) {
	v := os.Getenv(envvar)
	if v == "" {
		return "", ValueNotProvidedError
	}

	return v, nil
}

func (e *EnvironmentProvider) String(envvar string) StringProvider {
	return StringProviderFunc(func() (string, error) {
		v, err := e.get(envvar)
		if err != nil {
			return "", fmt.Errorf("EnvironmentProvider: %w", err)
		}

		return v, nil
	})
}

func (e *EnvironmentProvider) Int(envvar string) IntProvider {
	return IntProviderFunc(func() (int, error) {
		v, err := e.get(envvar)
		if err != nil {
			return 0, fmt.Errorf("EnvironmentProvider: %w", err)
		}

		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("EnvironmentProvider: failed to convert '%s' to int: %w", v, err)
		}

		return i, nil
	})
}
