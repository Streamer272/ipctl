package listener

import (
	"fmt"
	"github.com/Streamer272/ipcl/logger"
	"os"
	"os/exec"
	"strings"
	"time"
)

func didIpChange() (bool, string, error) {
	return true, "", nil
}

func Listen(interval int, command string, args []string) {
	log := logger.NewLogger()

	for {
		changed, newIp, err := didIpChange()
		if err != nil {
			log.Log("ERROR", fmt.Sprintf("Error occurred while listening to IP change, err = %v\n", err))
		}

		if changed {
			for i := 0; i < len(args); i++ {
				if strings.Contains(args[i], "$IP") {
					args[i] = strings.ReplaceAll(args[i], "$IP", newIp)
				}
			}
			command := exec.Command(command, args...)
			for i := 0; i < len(args); i++ {
				if strings.Contains(args[i], "$IP") {
					args[i] = strings.ReplaceAll(args[i], newIp, "$IP")
				}
			}

			command.Stdin = os.Stdin
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			err := command.Run()
			if err != nil {
				log.Log("ERROR", fmt.Sprintf("Error occurred while listening to IP change, err = %v\n", err))
			}

			log.Log("INFO", fmt.Sprintf("IP address changed"))
		}

		time.Sleep(time.Millisecond * time.Duration(interval))
	}
}
