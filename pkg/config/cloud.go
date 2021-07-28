package config

import "fmt"

// CloudConfig is the key-value store from the cloud name to it's URL
type CloudConfig map[string]*Cloud

// Cloud is a cloud object
type Cloud struct {
	URL          string `yaml:"url"`
	Bearer       string `yaml:"auth_token"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

// NewCloud creates a new Cloud object
func NewCloud(url, clientID, clientSecret, authToken string) *Cloud {
	return &Cloud{
		URL:          url,
		Bearer:       authToken,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

// NewCloudConfig creates a new cloud config object
func NewCloudConfig() CloudConfig {
	cloudConfig := make(CloudConfig)
	cloudConfig["imke"] = NewCloud("https://imke.cloud", "kubermatic", "", "")

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
