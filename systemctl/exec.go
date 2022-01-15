package systemctl

import (
	"github.com/Streamer272/ipctl/handle_error"
	"os"
	"os/exec"
)

func execSystemctl(command string) {
	cmd := exec.Command("systemctl", command, "ipctl.service")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	handle_error.HandleError(err)
}
