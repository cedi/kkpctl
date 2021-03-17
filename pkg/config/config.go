package config

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// ConfigPath is the path to our configuration file on disk
var ConfigPath string

// CloudConfig is the key-value store from the cloud name to it's URL
type CloudConfig map[string]Cloud

// Config is the configuration for KKPCTL
type Config struct {
	Provider ProviderConfig        `yaml:"provider"`
	Context  Context               `yaml:"ctx"`
	Cloud    CloudConfig           `yaml:"cloud"`
	OSSpec   OperatingSystemConfig `yaml:"os_spec"`
	NodeSpec CloudNodeConfig       `yaml:"node_spec"`
}

// NewConfig creates a new, empty, config
func NewConfig() Config {
	return Config{
		Provider: NewProvider(),
		Context:  NewContext(),
		Cloud:    make(map[string]Cloud),
	}
}

// Save saves a configuration
func (c *Config) Save() error {
	yamlByte, err := yaml.Marshal(c)
	if err != nil {
		return errors.Wrap(err, "Unable to serialize configuration")
	}

	err = ioutil.WriteFile(ConfigPath, yamlByte, 0600)
	if err != nil {
		return errors.Wrap(err, "Failed to write configuration")
	}

	return nil
}

// Read reads the config file and creates a empty config file if could not find a config file at the given path
func Read() (Config, error) {
	err := ensureConfig()
	if err != nil {
		return NewConfig(), errors.Wrap(err, "Failed to read kkpctl config")
	}

	data, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return NewConfig(), errors.Wrap(err, "Failed to read kkpctl config")
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return NewConfig(), errors.Wrap(err, "Failed to parse kkpctl config")
	}

	return config, nil
}

func ensureConfig() error {
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		path, _ := path.Split(ConfigPath)
		os.MkdirAll(path, 0640)

		config := NewConfig()
		err = config.Save()
		if err != nil {
			return errors.Wrap(err, "Failed to write empty kkpctl config")
		}
	}

	return nil
}

// GetCloudFromContext returns a touple of cloud-url and bearer
func (c *Config) GetCloudFromContext() (string, string) {
	return c.Cloud[c.Context.CloudName].URL, c.Cloud[c.Context.CloudName].Bearer
}
