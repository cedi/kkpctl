package config

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// Config is the configuration for KKPCTL
type Config struct {
	Provider ProviderConfig `yaml:"provider"`
	Context  Context        `yaml:"ctx"`
}

// NewConfig creates a new, empty, config
func NewConfig() Config {
	return Config{
		Provider: NewProvider(),
		Context:  NewContext(),
	}
}

// Save saves a configuration
func (c *Config) Save(path string) error {
	yamlByte, err := yaml.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "Unable to serialize configuration")
	}

	err = ioutil.WriteFile(path, yamlByte, 0600)
	if err != nil {
		return errors.Wrap(err, "Failed to write configuration")
	}

	return nil
}

// Read reads the config file and creates a empty config file if could not find a config file at the given path
func Read(path string) (Config, error) {
	ensureConfig(path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return NewConfig(), errors.Wrap(err, "Failed to read kkpctl config")
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return NewConfig(), errors.Wrap(err, "Failed to parse kkpctl config")
	}

	return config, nil
}

func ensureConfig(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		path, _ := path.Split(filePath)
		os.MkdirAll(path, 0640)

		config := NewConfig()
		err = config.Save(filePath)
		if err != nil {
			return errors.Wrap(err, "Failed to write empty kkpctl config")
		}
	}

	return nil
}
