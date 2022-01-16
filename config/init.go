package config

import (
	"encoding/json"
	"github.com/Streamer272/ipctl/constants"
	"github.com/Streamer272/ipctl/handle_error"
	"github.com/Streamer272/ipctl/helpers"
	"github.com/Streamer272/ipctl/options"
	"io/ioutil"
	"os"
)

func Init(rewrite bool) {
	if helpers.PathExists("/etc/ipctl/ipctl.json") && !rewrite {
		return
	}

	if !helpers.PathExists("/etc/ipctl") || rewrite {
		if helpers.PathExists("/etc/ipctl") && rewrite {
			err := os.RemoveAll("/etc/ipctl")
			handle_error.HandleError(err)
		}

		err := os.Mkdir("/etc/ipctl", constants.PERMS)
		handle_error.HandleError(err)
	}

	if !helpers.PathExists("/etc/ipctl/ipctl.json") || rewrite {
		optStr, err := json.Marshal(options.Default())
		handle_error.HandleError(err)

		err = ioutil.WriteFile("/etc/ipctl/ipctl.json", []byte(optStr), constants.PERMS)
		handle_error.HandleError(err)
	}
}

func Remove() {
	if helpers.PathExists("/etc/ipctl") {
		err := os.RemoveAll("/etc/ipctl")
		handle_error.HandleError(err)
	}
}
