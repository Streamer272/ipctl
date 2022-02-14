package config

import (
	"github.com/Streamer272/cool/check"
	"github.com/zpatrick/go-config"
)

var (
	EtcFile  = config.NewINIFile("/etc/ipctl/config")
	HomeFile = config.NewINIFile("$HOME/.config/ipctl/config")
)

func Get(name string) string {
	content := config.NewConfig([]config.Provider{EtcFile, HomeFile})
	err := content.Load()
	check.Check(err)

	value, err := content.String(name)
	check.Check(err)

	return value
}

func GetConfigFiles() []string {
	return []string{"/etc/ipctl/config", "$HOME/.config/ipctl/config"}
}
