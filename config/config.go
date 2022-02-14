package config

import (
	"github.com/Streamer272/cool/check"
	"github.com/zpatrick/go-config"
)

func Get(name string) string {
	etcFile := config.NewINIFile("/etc/ipctl/config")
	content := config.NewConfig([]config.Provider{etcFile})
	err := content.Load()
	check.Check(err)

	value, err := content.String(name)
	check.Check(err)

	return value
}

func GetConfigFiles() string {
	etcFile := config.NewINIFile("/etc/ipctl/config")
	content := config.NewConfig([]config.Provider{etcFile})
	err := content.Load()
	check.Check(err)

	output := ""
	for provider := range content.Providers {
		output += string(provider)
	}

	return output
}
