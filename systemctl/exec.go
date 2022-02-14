package systemctl

import (
	"os"
	"os/exec"
)

func Exec(command string, output bool) {
	cmd := exec.Command("systemctl", command, "ipctl.service")
	if output {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	cmd.Run()
}

func Logs() {
	cmd := exec.Command("journalctl", "-u", "ipctl.service")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
