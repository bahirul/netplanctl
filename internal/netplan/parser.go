package netplan

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadTemporaryOrRunConfig(temporary string, run string) (*NetplanConfig, error) {
	if _, err := os.Stat(temporary); err == nil {
		return LoadConfig(temporary)
	} else if os.IsNotExist(err) {
		return LoadConfig(run)
	} else {
		return nil, fmt.Errorf("failed to check temporary file: %w", err)
	}
}

func LoadConfig(filepath string) (*NetplanConfig, error) {
	data, err := os.ReadFile(filepath)

	if err != nil {
		return nil, fmt.Errorf("failed to read netplan file: %w", err)
	}

	var config NetplanConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml file: %w", err)
	}

	return &config, nil
}

func SaveConfig(config *NetplanConfig, filepath string) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal yaml file: %w", err)
	}

	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to write netplan file: %w", err)
	}
	return nil
}

func RawConfig(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Sprintf("failed to read netplan file: %v", err), err
	}
	return string(data), nil
}
