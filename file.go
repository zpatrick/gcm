package gcm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadJSONFile(path string, destination interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read '%s': %w", path, err)
	}

	if err := json.Unmarshal(content, destination); err != nil {
		return fmt.Errorf("failed to parse '%s': %w", path, err)
	}

	return nil
}

func LoadYAMLFile(path string, destination interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read '%s': %w", path, err)
	}

	if err := yaml.Unmarshal(content, destination); err != nil {
		return fmt.Errorf("failed to parse '%s': %w", path, err)
	}

	return nil
}
