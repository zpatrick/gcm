package gcm

import (
	"errors"
	"fmt"
	"reflect"
)

type Mux struct {
	Providers map[string]Provider
}

func (m *Mux) String(key string) (string, error) {
	provider, ok := m.Providers[key]
	if !ok {
		return "", fmt.Errorf("Mux: %s: %w", key, ProviderNotDefinedError)
	}

	v, err := provider.Value()
	if err != nil {
		return "", fmt.Errorf("Mux: %s: %w", key, err)
	}

	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("Mux: %s: %w", key, NewInvalidTypeError("string", reflect.TypeOf(v).String()))
	}

	return s, nil
}

func (m *Mux) MustString(key string) string {
	s, err := m.String(key)
	if err != nil {
		panic(err)
	}

	return s
}

func (m *Mux) Int(key string) (int, error) {
	return 0, nil
}

func (m *Mux) MustInt(key string) int {
	i, err := m.Int(key)
	if err != nil {
		panic(err)
	}

	return i
}

type Provider interface {
	Value() (interface{}, error)
}

type StringProvider interface {
	Value() (string, error)
}

type StringProviderFunc func() (string, error)

func (f StringProviderFunc) Value() (string, error) {
	return f()
}

type MultiStringProvider []StringProvider

func (m MultiStringProvider) Value() (string, error) {
	for _, p := range m {
		s, err := p.Value()
		if err != nil {
			if errors.Is(err, ValueNotProvidedError) {
				continue
			}

			return "", fmt.Errorf("MultiStringProvider: %w", err)
		}

		return s, nil
	}

	return "", fmt.Errorf("MultiStringProvider: %w", ValueNotProvidedError)
}

type StringValidator func(v string) error

type StringProviderSchema struct {
	Default       string
	DefaultIsZero bool
	Validate      StringValidator
	Provider      StringProvider
}

func (s *StringProviderSchema) Value() (interface{}, error) {
	v, err := s.Provider.Value()
	if err != nil {
		if !errors.Is(err, ValueNotProvidedError) {
			return nil, fmt.Errorf("StringProviderSchema: %w", err)
		}

		if s.Default == "" && !s.DefaultIsZero {
			return nil, fmt.Errorf("StringProviderSchema: %w", ValueNotProvidedError)
		}

		v = s.Default
	}

	if s.Validate != nil {
		if err := s.Validate(v); err != nil {
			return nil, fmt.Errorf("StringProviderSchema: %w", NewValidationError(err, v))
		}
	}

	return v, nil
}

type IntProvider interface {
	Value() (int, error)
}

type IntProviderFunc func() (int, error)

func (f IntProviderFunc) Value() (int, error) {
	return f()
}

type MultiIntProvider []IntProvider

func (m MultiIntProvider) Value() (int, error) {
	for _, p := range m {
		i, err := p.Value()
		if err != nil {
			if errors.Is(err, ValueNotProvidedError) {
				continue
			}

			return 0, fmt.Errorf("MultiIntProvider: %w", err)
		}

		return i, nil
	}

	return 0, fmt.Errorf("MultiIntProvider: %w", ValueNotProvidedError)
}

type IntValidator func(v int) error

func ValidateIntBetween(lower, upper int) IntValidator {
	return IntValidator(func(v int) error {
		switch {
		case v < lower:
			return fmt.Errorf("int %d is below lower limit %d", v, lower)
		case v > upper:
			return fmt.Errorf("int %d is above upper limit %d", v, upper)
		default:
			return nil
		}
	})
}

type IntProviderSchema struct {
	Default       int
	DefaultIsZero bool
	Validate      IntValidator
	Provider      IntProvider
}

func (s *IntProviderSchema) Value() (interface{}, error) {
	v, err := s.Provider.Value()
	if err != nil {
		if !errors.Is(err, ValueNotProvidedError) {
			return nil, fmt.Errorf("IntProviderSchema: %w", err)
		}

		if s.Default == 0 && !s.DefaultIsZero {
			return nil, fmt.Errorf("IntProviderSchema: %w", ValueNotProvidedError)
		}

		v = s.Default
	}

	if s.Validate != nil {
		if err := s.Validate(v); err != nil {
			return nil, fmt.Errorf("IntProviderSchema: %w", NewValidationError(err, v))
		}
	}

	return v, nil
}