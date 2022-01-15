package config

import (
	"encoding/json"
	"github.com/Streamer272/ipctl/handle_error"
	"github.com/Streamer272/ipctl/options"
	"io/ioutil"
	"os"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Init() {
	if exists("/etc/ipctl/ipctl.json") {
		return
	}

	if !exists("/etc/ipctl/ipctl.json") {
		err := os.Mkdir("/etc/ipctl", 744)
		handle_error.HandleError(err)
	}

	if !exists("/etc/ipctl/ipctl.json") {
		optStr, err := json.Marshal(options.Default())
		handle_error.HandleError(err)

		err = ioutil.WriteFile("/etc/ipctl/ipctl.json", []byte(optStr), 744)
		handle_error.HandleError(err)
	}

	// TODO: init systemctl service
}
