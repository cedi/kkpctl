package config

import (
	"testing"

	"github.com/go-test/deep"
)

func TestConfigParsing(t *testing.T) {
	config := New()
	config.Clouds = append(config.Clouds, Cloud{
		Name: "test",
		URL:  "imke.cloud",
	})

	config.BearerToken = "test123"
	config.Context = Context{
		Cloud:     "test",
		ProjectID: "x12vas",
	}
	config.ConfigPath = "../../tests/config/testfiles/config"

	config2, err := ReadFromConfig("../../tests/config/testfiles/config")
	if err != nil {
		t.Fatal(err.Error())
	}

	if diff := deep.Equal(config, config2); diff != nil {
		t.Error(diff)
	}
}
