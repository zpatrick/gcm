package main

import (
	"errors"
	"fmt"
)

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

			return 0, err
		}

		return i, nil
	}

	return 0, ValueNotProvidedError
}

func StaticInt(i int) IntProvider {
	return IntProviderFunc(func() (int, error) {
		return i, nil
	})
}

func OptionalStaticInt(i *int) IntProvider {
	if i != nil {
		return StaticInt(*i)
	}

	return IntProviderFunc(func() (int, error) {
		return 0, ValueNotProvidedError
	})
}

type IntValidator func(i int) error

type IntProviderSchema struct {
	Default       int
	DefaultIsZero bool
	Validate      IntValidator
	Provider      IntProvider
}

func (schema *IntProviderSchema) Value() (interface{}, error) {
	v, err := schema.Provider.Value()
	if err != nil {
		if !errors.Is(err, ValueNotProvidedError) {
			return nil, fmt.Errorf("IntProviderSchema: %w", err)
		}

		if schema.Default == 0 && !schema.DefaultIsZero {
			return nil, fmt.Errorf("IntProviderSchema: %w", ValueNotProvidedError)
		}

		v = schema.Default
	}

	if schema.Validate != nil {
		if err := schema.Validate(v); err != nil {
			return nil, fmt.Errorf("IntProviderSchema: %w", NewValidationError(err, v))
		}
	}

	return v, nil
}
