package gcm

// THIS IS AUTOGENERATED BLAH BLAH BLAH ...

import (
	"errors"
	"fmt"
	"time"
)

type Provider interface {
	Value() (interface{}, error)
}

type ProviderFunc func() (interface{}, error)

func (f ProviderFunc) Value() (interface{}, error) {
	return f()
}

type MultiProvider []Provider

func (m MultiProvider) Value() (interface{}, error) {
	for _, p := range m {
		v, err := p.Value()
		if err != nil {
			if errors.Is(err, ValueNotProvidedError) {
				continue
			}

			return nil, err
		}

		return v, nil
	}

	return nil, ValueNotProvidedError
}

func Static(v interface{}) Provider {
	return ProviderFunc(func() (interface{}, error) {
		return v, nil
	})
}

func OptionalStatic(v *interface{}) Provider {
	if v != nil {
		return Static(*v)
	}

	return ProviderFunc(func() (interface{}, error) {
		return nil, ValueNotProvidedError
	})
}

type Validator func(v interface{}) error

func ValidateInSet(values ...interface{}) Validator {
	set := make(map[interface{}]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return Validator(func(v interface{}) error {
		if _, ok := set[v]; !ok {
			printable := make([]string, 0, len(set))
			for v := range set {
				printable = append(printable, fmt.Sprintf("%v", v))
			}

			return fmt.Errorf("interface{} %v is not in set %v", v, printable)
		}

		return nil
	})
}

type ProviderSchema struct {
	Default       interface{}
	DefaultIsZero bool
	Validate      Validator
	Provider      Provider
}

func (schema *ProviderSchema) Value() (interface{}, error) {
	v, err := schema.Provider.Value()
	if err != nil {
		if !errors.Is(err, ValueNotProvidedError) {
			return nil, fmt.Errorf("ProviderSchema: %w", err)
		}

		if schema.Default == nil && !schema.DefaultIsZero {
			return nil, fmt.Errorf("ProviderSchema: %w", ValueNotProvidedError)
		}

		v = schema.Default
	}

	if schema.Validate != nil {
		if err := schema.Validate(v); err != nil {
			return nil, fmt.Errorf("ProviderSchema: %w", NewValidationError(err, v))
		}
	}

	return v, nil
}

type BoolProvider interface {
	Value() (bool, error)
}

type BoolProviderFunc func() (bool, error)

func (f BoolProviderFunc) Value() (bool, error) {
	return f()
}

type MultiBoolProvider []BoolProvider

func (m MultiBoolProvider) Value() (bool, error) {
	for _, p := range m {
		b, err := p.Value()
		if err != nil {
			if errors.Is(err, ValueNotProvidedError) {
				continue
			}

			return false, err
		}

		return b, nil
	}

	return false, ValueNotProvidedError
}

func StaticBool(b bool) BoolProvider {
	return BoolProviderFunc(func() (bool, error) {
		return b, nil
	})
}

func OptionalStaticBool(b *bool) BoolProvider {
	if b != nil {
		return StaticBool(*b)
	}

	return BoolProviderFunc(func() (bool, error) {
		return false, ValueNotProvidedError
	})
}

type BoolValidator func(b bool) error

type BoolProviderSchema struct {
	Default       bool
	DefaultIsZero bool
	Validate      BoolValidator
	Provider      BoolProvider
}

func (schema *BoolProviderSchema) Value() (interface{}, error) {
	v, err := schema.Provider.Value()
	if err != nil {
		if !errors.Is(err, ValueNotProvidedError) {
			return nil, fmt.Errorf("BoolProviderSchema: %w", err)
		}

		if schema.Default == false && !schema.DefaultIsZero {
			return nil, fmt.Errorf("BoolProviderSchema: %w", ValueNotProvidedError)
		}

		v = schema.Default
	}

	if schema.Validate != nil {
		if err := schema.Validate(v); err != nil {
			return nil, fmt.Errorf("BoolProviderSchema: %w", NewValidationError(err, v))
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

func ValidateIntInSet(values ...int) IntValidator {
	set := make(map[int]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return IntValidator(func(v int) error {
		if _, ok := set[v]; !ok {
			printable := make([]string, 0, len(set))
			for v := range set {
				printable = append(printable, fmt.Sprintf("%v", v))
			}

			return fmt.Errorf("int %v is not in set %v", v, printable)
		}

		return nil
	})
}

func ValidateIntBetween(lower, upper int) IntValidator {
	if lower > upper {
		panic(fmt.Sprintf("invalid validator: lower value %v is greater than upper value %v", lower, upper))
	}

	return IntValidator(func(v int) error {
		switch {
		case v < lower:
			return fmt.Errorf("%v is below lower limit %v", v, lower)
		case v > upper:
			return fmt.Errorf("%v is above upper limit %v", v, upper)
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

type Int64Provider interface {
	Value() (int64, error)
}

type Int64ProviderFunc func() (int64, error)

func (f Int64ProviderFunc) Value() (int64, error) {
	return f()
}

type MultiInt64Provider []Int64Provider

func (m MultiInt64Provider) Value() (int64, error) {
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

func StaticInt64(i int64) Int64Provider {
	return Int64ProviderFunc(func() (int64, error) {
		return i, nil
	})
}

func OptionalStaticInt64(i *int64) Int64Provider {
	if i != nil {
		return StaticInt64(*i)
	}

	return Int64ProviderFunc(func() (int64, error) {
		return 0, ValueNotProvidedError
	})
}

type Int64Validator func(i int64) error

func ValidateInt64InSet(values ...int64) Int64Validator {
	set := make(map[int64]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return Int64Validator(func(v int64) error {
		if _, ok := set[v]; !ok {
			printable := make([]string, 0, len(set))
			for v := range set {
				printable = append(printable, fmt.Sprintf("%v", v))
			}

			return fmt.Errorf("int64 %v is not in set %v", v, printable)
		}

		return nil
	})
}

func ValidateInt64Between(lower, upper int64) Int64Validator {
	if lower > upper {
		panic(fmt.Sprintf("invalid validator: lower value %v is greater than upper value %v", lower, upper))
	}

	return Int64Validator(func(v int64) error {
		switch {
		case v < lower:
			return fmt.Errorf("%v is below lower limit %v", v, lower)
		case v > upper:
			return fmt.Errorf("%v is above upper limit %v", v, upper)
		default:
			return nil
		}
	})
}

type Int64ProviderSchema struct {
	Default       int64
	DefaultIsZero bool
	Validate      Int64Validator
	Provider      Int64Provider
}

func (schema *Int64ProviderSchema) Value() (interface{}, error) {
	v, err := schema.Provider.Value()
	if err != nil {
		if !errors.Is(err, ValueNotProvidedError) {
			return nil, fmt.Errorf("Int64ProviderSchema: %w", err)
		}

		if schema.Default == 0 && !schema.DefaultIsZero {
			return nil, fmt.Errorf("Int64ProviderSchema: %w", ValueNotProvidedError)
		}

		v = schema.Default
	}

	if schema.Validate != nil {
		if err := schema.Validate(v); err != nil {
			return nil, fmt.Errorf("Int64ProviderSchema: %w", NewValidationError(err, v))
		}
	}

	return v, nil
}

type Float64Provider interface {
	Value() (float64, error)
}

type Float64ProviderFunc func() (float64, error)

func (f Float64ProviderFunc) Value() (float64, error) {
	return f()
}

type MultiFloat64Provider []Float64Provider

func (m MultiFloat64Provider) Value() (float64, error) {
	for _, p := range m {
		f, err := p.Value()
		if err != nil {
			if errors.Is(err, ValueNotProvidedError) {
				continue
			}

			return 0, err
		}

		return f, nil
	}

	return 0, ValueNotProvidedError
}

func StaticFloat64(f float64) Float64Provider {
	return Float64ProviderFunc(func() (float64, error) {
		return f, nil
	})
}

func OptionalStaticFloat64(f *float64) Float64Provider {
	if f != nil {
		return StaticFloat64(*f)
	}

	return Float64ProviderFunc(func() (float64, error) {
		return 0, ValueNotProvidedError
	})
}

type Float64Validator func(f float64) error

func ValidateFloat64InSet(values ...float64) Float64Validator {
	set := make(map[float64]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return Float64Validator(func(v float64) error {
		if _, ok := set[v]; !ok {
			printable := make([]string, 0, len(set))
			for v := range set {
				printable = append(printable, fmt.Sprintf("%v", v))
			}

			return fmt.Errorf("float64 %v is not in set %v", v, printable)
		}

		return nil
	})
}

func ValidateFloat64Between(lower, upper float64) Float64Validator {
	if lower > upper {
		panic(fmt.Sprintf("invalid validator: lower value %v is greater than upper value %v", lower, upper))
	}

	return Float64Validator(func(v float64) error {
		switch {
		case v < lower:
			return fmt.Errorf("%v is below lower limit %v", v, lower)
		case v > upper:
			return fmt.Errorf("%v is above upper limit %v", v, upper)
		default:
			return nil
		}
	})
}

type Float64ProviderSchema struct {
	Default       float64
	DefaultIsZero bool
	Validate      Float64Validator
	Provider      Float64Provider
}

func (schema *Float64ProviderSchema) Value() (interface{}, error) {
	v, err := schema.Provider.Value()
	if err != nil {
		if !errors.Is(err, ValueNotProvidedError) {
			return nil, fmt.Errorf("Float64ProviderSchema: %w", err)
		}

		if schema.Default == 0 && !schema.DefaultIsZero {
			return nil, fmt.Errorf("Float64ProviderSchema: %w", ValueNotProvidedError)
		}

		v = schema.Default
	}

	if schema.Validate != nil {
		if err := schema.Validate(v); err != nil {
			return nil, fmt.Errorf("Float64ProviderSchema: %w", NewValidationError(err, v))
		}
	}

	return v, nil
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

			return "", err
		}

		return s, nil
	}

	return "", ValueNotProvidedError
}

func StaticString(s string) StringProvider {
	return StringProviderFunc(func() (string, error) {
		return s, nil
	})
}

func OptionalStaticString(s *string) StringProvider {
	if s != nil {
		return StaticString(*s)
	}

	return StringProviderFunc(func() (string, error) {
		return "", ValueNotProvidedError
	})
}

type StringValidator func(s string) error

func ValidateStringInSet(values ...string) StringValidator {
	set := make(map[string]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return StringValidator(func(v string) error {
		if _, ok := set[v]; !ok {
			printable := make([]string, 0, len(set))
			for v := range set {
				printable = append(printable, fmt.Sprintf("%v", v))
			}

			return fmt.Errorf("string %v is not in set %v", v, printable)
		}

		return nil
	})
}

type StringProviderSchema struct {
	Default       string
	DefaultIsZero bool
	Validate      StringValidator
	Provider      StringProvider
}

func (schema *StringProviderSchema) Value() (interface{}, error) {
	v, err := schema.Provider.Value()
	if err != nil {
		if !errors.Is(err, ValueNotProvidedError) {
			return nil, fmt.Errorf("StringProviderSchema: %w", err)
		}

		if schema.Default == "" && !schema.DefaultIsZero {
			return nil, fmt.Errorf("StringProviderSchema: %w", ValueNotProvidedError)
		}

		v = schema.Default
	}

	if schema.Validate != nil {
		if err := schema.Validate(v); err != nil {
			return nil, fmt.Errorf("StringProviderSchema: %w", NewValidationError(err, v))
		}
	}

	return v, nil
}

type DurationProvider interface {
	Value() (time.Duration, error)
}

type DurationProviderFunc func() (time.Duration, error)

func (f DurationProviderFunc) Value() (time.Duration, error) {
	return f()
}

type MultiDurationProvider []DurationProvider

func (m MultiDurationProvider) Value() (time.Duration, error) {
	for _, p := range m {
		d, err := p.Value()
		if err != nil {
			if errors.Is(err, ValueNotProvidedError) {
				continue
			}

			return 0, err
		}

		return d, nil
	}

	return 0, ValueNotProvidedError
}

func StaticDuration(d time.Duration) DurationProvider {
	return DurationProviderFunc(func() (time.Duration, error) {
		return d, nil
	})
}

func OptionalStaticDuration(d *time.Duration) DurationProvider {
	if d != nil {
		return StaticDuration(*d)
	}

	return DurationProviderFunc(func() (time.Duration, error) {
		return 0, ValueNotProvidedError
	})
}

type DurationValidator func(d time.Duration) error

func ValidateDurationInSet(values ...time.Duration) DurationValidator {
	set := make(map[time.Duration]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return DurationValidator(func(v time.Duration) error {
		if _, ok := set[v]; !ok {
			printable := make([]string, 0, len(set))
			for v := range set {
				printable = append(printable, fmt.Sprintf("%v", v))
			}

			return fmt.Errorf("time.Duration %v is not in set %v", v, printable)
		}

		return nil
	})
}

func ValidateDurationBetween(lower, upper time.Duration) DurationValidator {
	if lower > upper {
		panic(fmt.Sprintf("invalid validator: lower value %v is greater than upper value %v", lower, upper))
	}

	return DurationValidator(func(v time.Duration) error {
		switch {
		case v < lower:
			return fmt.Errorf("%v is below lower limit %v", v, lower)
		case v > upper:
			return fmt.Errorf("%v is above upper limit %v", v, upper)
		default:
			return nil
		}
	})
}

type DurationProviderSchema struct {
	Default       time.Duration
	DefaultIsZero bool
	Validate      DurationValidator
	Provider      DurationProvider
}

func (schema *DurationProviderSchema) Value() (interface{}, error) {
	v, err := schema.Provider.Value()
	if err != nil {
		if !errors.Is(err, ValueNotProvidedError) {
			return nil, fmt.Errorf("DurationProviderSchema: %w", err)
		}

		if schema.Default == 0 && !schema.DefaultIsZero {
			return nil, fmt.Errorf("DurationProviderSchema: %w", ValueNotProvidedError)
		}

		v = schema.Default
	}

	if schema.Validate != nil {
		if err := schema.Validate(v); err != nil {
			return nil, fmt.Errorf("DurationProviderSchema: %w", NewValidationError(err, v))
		}
	}

	return v, nil
}
