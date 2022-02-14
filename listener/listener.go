package listener

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/Streamer272/cool/check"
)

var ip string = ""

func GetCurrentIp() (string, error) {
	response, err := http.Get("https://api.my-ip.io/ip.json")
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var data struct {
		Success bool   `json:"success"`
		IP      string `json:"ip"`
		Type    string `json:"type"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if data.Success != true {
		return "", errors.New("request not successful")
	}

	return data.IP, nil
}

func Listen(command string, interval int) {
	//log := logger.NewLogger()

	//log.Log("INFO", fmt.Sprintf("Listening"))

	defer func() {
		if err := recover(); err != nil {
			//log.Log("ERROR", fmt.Sprintf("Error occurred while listening to IP change (%v)\n", err))
			Listen(command, interval)
		}
	}()

	for {
		newIp, err := GetCurrentIp()
		if err != nil {
			//log.Log("ERROR", fmt.Sprintf("Error occurred while retrieving current ip (%v)\n", err))
		}

		fmt.Printf("%v == %v\n", newIp, ip)

		if newIp == ip {
			time.Sleep(time.Millisecond * time.Duration(interval))
			continue
		}

		if ip == "" {
			ip = newIp
			time.Sleep(time.Millisecond * time.Duration(interval))
			continue
		}

		ip = newIp

		err = os.Setenv("IP", ip)
		if err != nil {
			//log.Log("ERROR", fmt.Sprintf("Error occurred while chaning ENV (%v)\n", err))
		}

		command := exec.Command("/usr/bin/bash", command)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		err = command.Run()
		if err != nil {
			//log.Log("ERROR", fmt.Sprintf("Error occurred while running command (%v)\n", err))
		}

		//log.Log("INFO", fmt.Sprintf("IP address changed"))

		time.Sleep(time.Millisecond * time.Duration(interval))
	}
}

func Update(command string) {
	//log := logger.NewLogger()

	//log.Log("INFO", fmt.Sprintf("Updating DNS records"))

	ip, err := GetCurrentIp()
	check.Check(err)
	if ip == "" {
		//log.Log("ERROR", fmt.Sprintf("IP (%v) is undefined", ip))
	}

	err = os.Setenv("IP", ip)
	if err != nil {
		//log.Log("ERROR", fmt.Sprintf("Error occurred while changing ENV (%v)\n", ip))
	}

	execCommand := exec.Command("/usr/bin/bash", command)
	execCommand.Stdin = os.Stdin
	execCommand.Stdout = os.Stdout
	execCommand.Stderr = os.Stderr
	err = execCommand.Run()
	if err != nil {
		//log.Log("ERROR", fmt.Sprintf("Error occurred while running command (%v)\n", err))
	}

	//log.Log("INFO", fmt.Sprintf("DNS records updated"))
}
