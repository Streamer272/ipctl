package listener

import (
	"encoding/json"
	"fmt"
	"github.com/Streamer272/ipcl/logger"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// TODO: save me to file
var ip string = ""

type Response struct {
	Success bool   `json:"Success"`
	IP      string `json:"ip"`
	Type    string `json:"type"`
}

type RequestNotSuccessful struct {
	response Response
}

func (err RequestNotSuccessful) Error() string {
	return err.response.IP
}

func getCurrentIp() (string, error) {
	response, err := http.Get("https://api.my-ip.io/ip.json")
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if data.Success != true {
		return "", RequestNotSuccessful{response: data}
	}

	return data.IP, nil
}

func didIpChange() (bool, error) {
	currentIp, err := getCurrentIp()
	if err != nil {
		return false, err
	}

	if ip == "" {
		ip = currentIp
		return false, nil
	}
	if ip != currentIp {
		ip = currentIp
		return true, nil
	}

	return false, nil
}

func Listen(interval int, command string, args []string) {
	log := logger.NewLogger()

	for {
		changed, err := didIpChange()
		if err != nil {
			log.Log("ERROR", fmt.Sprintf("Error occurred while listening to IP change, err = %v\n", err))
		}

		if changed {
			for i := 0; i < len(args); i++ {
				if strings.Contains(args[i], "$IP") {
					args[i] = strings.ReplaceAll(args[i], "$IP", ip)
				}
			}
			command := exec.Command(command, args...)
			for i := 0; i < len(args); i++ {
				if strings.Contains(args[i], "$IP") {
					args[i] = strings.ReplaceAll(args[i], ip, "$IP")
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
