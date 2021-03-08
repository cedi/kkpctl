package main

import (
	"fmt"

	kkpconfig "github.com/cedi/kkpctl/pkg/config"
)

func main() {

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

	config.Save()

	config2, err := kkpconfig.ReadFromConfig("/Users/cedi/.config/kkpctl/config")
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	fmt.Printf("%v\n", config2)
}
