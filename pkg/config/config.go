package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/cedi/kkpctl/pkg/errors"
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

// NewCloudConfig creates a new cloud config object
func NewCloudConfig() CloudConfig {
	return make(CloudConfig)
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
			return errors.Wrap(err, "failed to write empty kkpctl config")
		}
	}

	return nil
}

// GetCloudFromContext returns a touple of cloud-url and bearer
func (c *Config) getCloudFromContext() (string, string) {
	return c.Cloud[c.Context.CloudName].URL, c.Cloud[c.Context.CloudName].Bearer
}

// GetKKPClient returns a KKP Client for the currently configured cloud
func (c *Config) GetKKPClient() (*client.Client, error) {
	baseUrl, apiToken := c.getCloudFromContext()
	kkp, err := client.NewClient(baseUrl, apiToken)
	if err != nil {
		return kkp, errors.Wrap(err, "could not initialize Kubermatic API client")
	}

	return kkp, nil
}

// Set sets a cloud
func (c *CloudConfig) Set(name string, value Cloud) {
	(*c)[name] = value
}

// Get gets a cloud
func (c *CloudConfig) Get(name string) (*Cloud, error) {
	cloud, ok := (*c)[name]
	if !ok {
		return nil, fmt.Errorf("cannot find cloud %s in your configuration", name)
	}
	return &cloud, nil
}
