package systemctl

import (
	"github.com/Streamer272/ipctl/constants"
	"github.com/Streamer272/ipctl/handle_error"
	"github.com/Streamer272/ipctl/helpers"
	"io/ioutil"
	"os"
	"os/exec"
)

func Init(enable bool) {
	if !helpers.PathExists("/lib/systemd/system/ipctl.service") {
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

		if enable {
			command := exec.Command("systemctl", "enable", "ipctl.service")
			command.Stdin = os.Stdin
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			err = command.Run()
			handle_error.HandleError(err)
		}
	}
}

func Remove() {
	if helpers.PathExists("/lib/systemd/system/ipctl.service") {
		Disable()
		err := os.RemoveAll("/lib/systemd/system/ipctl.service")
		handle_error.HandleError(err)
	}
}
