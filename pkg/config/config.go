package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/cedi/kkpctl/pkg/client"
	"github.com/phayes/permbits"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// ConfigPath is the path to our configuration file on disk
var ConfigPath string

// Config is the configuration for KKPCTL
type Config struct {
	Provider *ProviderConfig        `yaml:"provider"`
	Context  *Context               `yaml:"ctx"`
	Cloud    CloudConfig            `yaml:"cloud"`
	OSSpec   *OperatingSystemConfig `yaml:"os_spec"`
	NodeSpec *CloudNodeConfig       `yaml:"node_spec"`
}

// NewConfig creates a new, empty, config
func NewConfig() *Config {
	return &Config{
		Provider: NewProvider(),
		Context:  NewContext(),
		Cloud:    NewCloudConfig(),
		OSSpec:   NewOSSpecConfig(),
		NodeSpec: NewNodeSpecConfig(),
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
func Read() (*Config, error) {
	err := EnsureConfig()
	if err != nil {
		return NewConfig(), errors.Wrap(err, "Failed to read kkpctl config")
	}

	data, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return NewConfig(), errors.Wrap(err, "Failed to read kkpctl config")
	}

	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return NewConfig(), errors.Wrap(err, "Failed to parse kkpctl config")
	}

	// fix a bug where cluster creation does not work when the URL contains a trailing slash
	for _, cloud := range config.Cloud {
		if cloud == nil {
			continue
		}

		cloud.URL = strings.TrimSuffix(cloud.URL, "/")
	}

	return config, nil
}

// EnsureConfig ensures that the config is there and exists using the correct permissions.
func EnsureConfig() error {
	// Create config file if not exists
	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		path, _ := path.Split(ConfigPath)
		os.MkdirAll(path, 0600)

		config := NewConfig()
		err = config.Save()
		if err != nil {
			return errors.Wrap(err, "failed to write empty kkpctl config")
		}
	}

	return checkPermissions(ConfigPath)
}

func checkPermissions(configPath string) error {
	permissions, err := permbits.Stat(configPath)
	if err != nil {
		return errors.Wrap(err, "failed to check file permissions for kkpctl config")
	}

	if permissions.UserExecute() {
		fmt.Println("")
		fmt.Println("!! Configuration anomaly detected")
		fmt.Println("!! kkpctl config should not be executable")
		fmt.Println("!! to get rid of this message, run")
		fmt.Printf("!! $ chmod -R 600 %s\n", configPath)
		fmt.Println("")
	}

	if permissions.GroupRead() || permissions.GroupWrite() || permissions.GroupExecute() ||
		permissions.OtherRead() || permissions.OtherWrite() || permissions.OtherExecute() {
		fmt.Println("")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("!! Security Alert")
		fmt.Println("!! kkpctl config is world readable!")
		fmt.Println("!! Since kkpctl contains passwords and OAuth Tokens, this could cause serious security issues!")
		fmt.Println("!! to get rid of this message, run")
		fmt.Printf("!! $ chmod -R 600 %s\n", configPath)
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("")
	}
	return nil
}

// GetCloudFromContext returns a touple of cloud-url and bearer
func (c *Config) getCloudFromContext() (string, string) {
	if c.Context == nil {
		return "", ""
	}

	cloud, ok := c.Cloud[c.Context.CloudName]
	if !ok {
		return "", ""
	}

	if cloud == nil {
		return "", ""
	}

	return cloud.URL, cloud.Bearer
}

// GetKKPClient returns a KKP Client for the currently configured cloud
func (c *Config) GetKKPClient(apiVersion client.APIVersion) (*client.Client, error) {
	baseURL, apiToken := c.getCloudFromContext()
	kkp, err := client.NewClient(apiVersion, baseURL, apiToken)
	if err != nil {
		return kkp, errors.Wrap(err, "could not initialize Kubermatic API client")
	}

	return kkp, nil
}
