package config

import (
	"io/ioutil"
	"os"

	"github.com/Streamer272/cool/check"
	"github.com/Streamer272/ipctl/constants"
)

func Init() {
	if _, err := os.Stat("/etc/ipctl"); err != nil {
		err := os.Mkdir("/etc/ipctl", constants.PERMS)
		check.Check(err)
	}

	err := ioutil.WriteFile("/etc/ipctl/config", []byte(constants.DEFAULT_CONFIG), constants.PERMS)
	check.Check(err)
}
