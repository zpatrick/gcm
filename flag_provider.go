package gcm

import (
	"flag"
	"fmt"
	"os"
)

type FlagProvider struct {
	*flag.FlagSet
}

func NewFlagProvider(f *flag.FlagSet) *FlagProvider {
	return &FlagProvider{f}
}

func (f *FlagProvider) assertParsed() error {
	if !f.Parsed() {
		args := os.Args[1:]
		if err := f.Parse(args); err != nil {
			return fmt.Errorf("failed to parse args %v: %w", args, err)
		}
	}

	return nil
}

func (f *FlagProvider) isFlagPassed(name string) bool {
	var passed bool
	f.Visit(func(f *flag.Flag) {
		if f.Name == name {
			passed = true
		}
	})

	return passed
}

func (f *FlagProvider) Int(name string, value int, usage string, useDefault bool) IntProvider {
	ptr := f.FlagSet.Int(name, value, usage)
	return IntProviderFunc(func() (int, error) {
		if err := f.assertParsed(); err != nil {
			return 0, fmt.Errorf("FlagProvider: %w", err)
		}

		if !useDefault && !f.isFlagPassed(name) {
			return 0, fmt.Errorf("FlagProvider: %w", ValueNotProvidedError)
		}

		return *ptr, nil
	})
}

func (f *FlagProvider) String(name, value, usage string, useDefault bool) StringProvider {
	ptr := f.FlagSet.String(name, value, usage)
	return StringProviderFunc(func() (string, error) {
		if err := f.assertParsed(); err != nil {
			return "", fmt.Errorf("FlagProvider: %w", err)
		}

		if !useDefault && !f.isFlagPassed(name) {
			return "", fmt.Errorf("FlagProvider: %w", ValueNotProvidedError)
		}

		return *ptr, nil
	})
}
