package systemctl

import (
	"os"
	"os/exec"
)

func execSystemctl(command string) {
	cmd := exec.Command("systemctl", command, "ipctl.service")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//goland:noinspection GoUnhandledErrorResult
	cmd.Run()
}
