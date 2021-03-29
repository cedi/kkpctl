package config

import "fmt"

// CloudConfig is the key-value store from the cloud name to it's URL
type CloudConfig map[string]*Cloud

// Cloud is a cloud object
type Cloud struct {
	URL    string `yaml:"url"`
	Bearer string `yaml:"bearer"`
}

// NewCloud creates a new Cloud object
func NewCloud(url, bearer string) *Cloud {
	return &Cloud{
		URL:    url,
		Bearer: bearer,
	}
}

// NewCloudConfig creates a new cloud config object
func NewCloudConfig() CloudConfig {
	cloudConfig := make(CloudConfig)
	cloudConfig["imke"] = NewCloud("https://imke.cloud", "")

	return cloudConfig
}

// Set sets a cloud
func (c *CloudConfig) Set(name string, value *Cloud) {
	(*c)[name] = value
}

// Get gets a cloud
func (c *CloudConfig) Get(name string) (*Cloud, error) {
	cloud, ok := (*c)[name]
	if !ok {
		return nil, fmt.Errorf("cannot find cloud %s in your configuration", name)
	}
	return cloud, nil
}
