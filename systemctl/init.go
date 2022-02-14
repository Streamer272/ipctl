package systemctl

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/Streamer272/cool/check"
	"github.com/Streamer272/ipctl/constants"
)

func Init(enable bool) {
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
	check.Check(err)

	if enable {
		command := exec.Command("systemctl", "enable", "ipctl.service")
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Run()
	}
}
