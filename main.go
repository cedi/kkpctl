package main

import (
	"fmt"

	kkpconfig "github.com/cedi/kkpctl/pkg/config"
)

// Config is the global configuration for the kkpctl
var Config *kkpconfig.Config

func main() {
	var err error
	Config, err = kkpconfig.ReadFromConfig("/Users/cedi/.config/kkpctl/config")
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
}
