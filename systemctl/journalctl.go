package systemctl

import (
	"os"
	"os/exec"
)

func Logs() {
	cmd := exec.Command("journalctl", "-u", "ipctl.service")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//goland:noinspection GoUnhandledErrorResult
	cmd.Run()
}
