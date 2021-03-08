package config

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// Cloud contains the Name and URL to a KKP Installation
type Cloud struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

// Context is the live context to be used in kkpctl
type Context struct {
	Cloud     string `yaml:"cloud"`
	ProjectID string `yaml:"projectID"`
}

// Config is serialized to ~/.config/kkpctl/config and contains the global configuration for kkpctl
type Config struct {
	Clouds      []Cloud `yaml:"clouds"`
	Context     Context `yaml:"context"`
	BearerToken string  `yaml:"bearer"`
	ConfigPath  string  `yaml:",omitempty"`
}

// New creates an empty configuration
func New() *Config {
	return &Config{
		Clouds:      make([]Cloud, 0),
		Context:     Context{},
		BearerToken: "",
		ConfigPath:  "",
	}
}

// NewWithBearer creates a new config with a bearer token
func NewWithBearer(bearer string) *Config {
	config := New()
	config.BearerToken = bearer
	return config
}

// ReadFromConfig reads the config file
func ReadFromConfig(location string) (*Config, error) {
	data, err := ioutil.ReadFile(location)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read kkpctl config")
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return &config, errors.Wrap(err, "Failed to parse kkpctl config")
	}

	config.ConfigPath = location

	return &config, nil
}

// Save saves a configuration
func Save(location string, config *Config) error {
	if location != config.ConfigPath {
		return fmt.Errorf("Configuration contains a different path than the location argument. Consider using config.Save()")
	}

	config.ConfigPath = ""
	yaml, err := yaml.Marshal(config)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal configuration")
	}
	config.ConfigPath = location

	err = ioutil.WriteFile(location, yaml, 0644)
	if err != nil {
		return errors.Wrap(err, "Failed to write configuration")
	}

	return nil
}

// Save saves a configuration
func (c *Config) Save() error {
	return Save(c.ConfigPath, c)
}
