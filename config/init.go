package config

import (
	"encoding/json"
	"github.com/Streamer272/ipctl/constants"
	"github.com/Streamer272/ipctl/handle_error"
	"github.com/Streamer272/ipctl/options"
	"io/ioutil"
	"os"
	"os/exec"
)

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Init(dontEnable bool, force bool) {
	if pathExists("/etc/ipctl/ipctl.json") && !force {
		return
	}

	if !pathExists("/etc/ipctl") || force {
		if pathExists("/etc/ipctl") && force {
			err := os.RemoveAll("/etc/ipctl")
			handle_error.HandleError(err)
		}

		err := os.Mkdir("/etc/ipctl", constants.PERMS)
		handle_error.HandleError(err)
	}

	if !pathExists("/etc/ipctl/ipctl.json") || force {
		optStr, err := json.Marshal(options.Default())
		handle_error.HandleError(err)

		err = ioutil.WriteFile("/etc/ipctl/ipctl.json", []byte(optStr), constants.PERMS)
		handle_error.HandleError(err)
	}

	if !pathExists("/lib/systemd/system/ipctl.service") || force {
		err := ioutil.WriteFile("/lib/systemd/system/ipctl.service", []byte(""+
			"[Unit]\n"+
			"Description=Listen to IP change and change your DNS' records dynamically\n"+
			"After=network.target\n"+
			"StartLimitIntervalSec=0\n\n"+
			"[Service]\n"+
			"Type=simple\n"+
			"Restart=always\n"+
			"RestartSec=1\n"+
			"User=root\n"+
			"ExecStart=ipctl listen\n\n"+
			"[Install]\n"+
			"WantedBy=multi-currentUser.target\n",
		), constants.PERMS)
		handle_error.HandleError(err)

		if !dontEnable {
			command := exec.Command("systemctl", "enable", "ipctl.service")
			command.Stdin = os.Stdin
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			err = command.Run()
			handle_error.HandleError(err)
		}
	}
}
