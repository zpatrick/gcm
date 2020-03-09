package gcm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"gopkg.in/yaml.v2"
)

type DataParser func(content []byte) (fields map[string]interface{}, err error)

func YAMLParser(content []byte) (map[string]interface{}, error) {
	var fields map[string]interface{}
	if err := yaml.Unmarshal(content, &fields); err != nil {
		return nil, fmt.Errorf("failed to parse yaml: %w", err)
	}

	return fields, nil
}

func JSONParser(content []byte) (map[string]interface{}, error) {
	var fields map[string]interface{}
	if err := json.Unmarshal(content, &fields); err != nil {
		return nil, fmt.Errorf("failed to parse json: %w", err)
	}

	return fields, nil
}

type ReloadPolicy func() bool

func ReloadAlways() bool { return true }

func ReloadNever() bool { return false }

type FileProvider struct {
	Path   string
	Parser DataParser
	Reload ReloadPolicy
	fields map[string]interface{}
}

func NewFileProvider(path string, parser DataParser, reload ReloadPolicy) FileProvider {
	return FileProvider{
		Path:   path,
		Parser: parser,
		Reload: reload,
	}
}

func (f *FileProvider) loadFields() (map[string]interface{}, error) {
	if f.fields == nil || f.Reload() {
		content, err := ioutil.ReadFile(f.Path)
		if err != nil {
			return nil, fmt.Errorf("failed to read '%s': %w", f.Path, err)
		}

		fields, err := f.Parser(content)
		if err != nil {
			return nil, fmt.Errorf("failed to parse file contents: %w", err)
		}

		f.fields = fields
	}

	return f.fields, nil
}

func (f *FileProvider) loadField(section string, subsections ...string) (interface{}, error) {
	fields, err := f.loadFields()
	if err != nil {
		return nil, err
	}

	chain := append([]string{section}, subsections...)
	for ; len(chain) > 1; chain = chain[1:] {
		sectionKey := chain[0]

		section, ok := fields[sectionKey]
		if !ok {
			return nil, fmt.Errorf("section '%s' not defined: %w", sectionKey, ValueNotProvidedError)
		}

		switch sec := section.(type) {
		case map[string]interface{}:
			fields = sec
		case map[interface{}]interface{}:
			sectionFields := make(map[string]interface{}, len(sec))
			for k, v := range sec {
				ks, ok := k.(string)
				if !ok {
					ite := NewInvalidTypeError("string", reflect.TypeOf(k).String())
					return nil, fmt.Errorf("section '%s' contained a non-string key: %w", sectionKey, ite)
				}

				sectionFields[ks] = v
			}

			fields = sectionFields
		default:
			ite := NewInvalidTypeError("map", reflect.TypeOf(section).String())
			return nil, fmt.Errorf("section '%s' is not a map type: %w", sectionKey, ite)
		}
	}

	val, ok := fields[chain[0]]
	if !ok {
		return nil, ValueNotProvidedError
	}

	return val, nil
}

func (f *FileProvider) String(section string, subsections ...string) StringProvider {
	return StringProviderFunc(func() (string, error) {
		v, err := f.loadField(section, subsections...)
		if err != nil {
			return "", fmt.Errorf("FileProvider: %w", err)
		}

		s, ok := v.(string)
		if !ok {
			return "", fmt.Errorf("FileProvder: %w", NewInvalidTypeError("string", reflect.TypeOf(v).String()))
		}

		return s, nil
	})
}

func (f *FileProvider) Int(section string, subsections ...string) IntProvider {
	return IntProviderFunc(func() (int, error) {
		v, err := f.loadField(section, subsections...)
		if err != nil {
			return 0, err
		}

		i, ok := v.(int)
		if !ok {
			return 0, fmt.Errorf("FileProvder: %w", NewInvalidTypeError("int", reflect.TypeOf(v).String()))
		}

		return i, nil
	})
}
