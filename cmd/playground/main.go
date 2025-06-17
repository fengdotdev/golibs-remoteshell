package main

import (
	"fmt"

	"github.com/fengdotdev/golibs-remoteshell/sandbox/draft1/settings"
	"github.com/fengdotdev/golibs-remoteshell/sandbox/draft1/shell"
)

func main() {
	shell.Terminal()
}

func Init() {

	config := settings.New("mykey", "myuser", "mypassword")

	err := settings.SaveSettings(config)
	if err != nil {
		panic(err)
	}

	config2, err := settings.GetSettingsFrom()
	if err != nil {
		panic(err)
	}
	fmt.Println("Config2:", config2.GetKey())
}
