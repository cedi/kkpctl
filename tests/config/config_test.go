package main

import (
	"testing"

	kkpconfig "github.com/cedi/kkpctl/pkg/config"
	"github.com/go-test/deep"
)

func TestConfigParsing(t *testing.T) {
	config := kkpconfig.New()
	config.Clouds = append(config.Clouds, kkpconfig.Cloud{
		Name: "test",
		URL:  "imke.cloud",
	})

	config.BearerToken = "test123"
	config.Context = kkpconfig.Context{
		Cloud:     "test",
		ProjectID: "x12vas",
	}
	config.ConfigPath = "testfiles/config"

	config2, err := kkpconfig.ReadFromConfig("testfiles/config")
	if err != nil {
		t.Fatal(err.Error())
	}

	if diff := deep.Equal(config, config2); diff != nil {
		t.Error(diff)
	}
}
