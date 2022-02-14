package listener

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var currentIp = ""

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
	log.Printf("Listening to IP change\n")

	defer func() {
		if err := recover(); err != nil {
			log.Printf("Fatal error occurred: %v\n", err)
			Listen(command, interval)
		}
	}()

	for {
		ip, err := GetCurrentIp()
		if err != nil || ip == "" {
			log.Printf("Error occurred while fetching IP: %v\n", err)
			continue
		}

		log.Printf("Current IP: %v\n", ip)

		if ip == currentIp {
			time.Sleep(time.Millisecond * time.Duration(interval))
			continue
		}

		if currentIp == "" {
			currentIp = ip
			time.Sleep(time.Millisecond * time.Duration(interval))
			continue
		}

		currentIp = ip

		err = os.Setenv("IP", currentIp)
		if err != nil {
			log.Printf("Error occurred while changing ENV: %v -> Cannot execute callback file\n", err)
			continue
		}

		command := exec.Command("/usr/bin/bash", command)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		err = command.Run()
		if err != nil {
			log.Printf("Error occurred while executing callback file: %v\n", err)
			time.Sleep(time.Millisecond * time.Duration(interval))
			continue
		}

		log.Printf("IP address changed")
		time.Sleep(time.Millisecond * time.Duration(interval))
	}
}

func Update(command string) {
	log.Printf("Updating DNS' records, sit tight...")

	ip, err := GetCurrentIp()
	if err != nil || ip == "" {
		//log.Log("ERROR", fmt.Sprintf("IP (%v) is undefined", ip))
		log.Printf("Error occurred while fetching IP: %v\n", err)
		return
	}

	err = os.Setenv("IP", ip)
	if err != nil {
		log.Printf("Error occurred while changing ENV: %v -> Cannot execute callback file\n", err)
		return
	}

	execCommand := exec.Command("/usr/bin/bash", command)
	execCommand.Stdin = os.Stdin
	execCommand.Stdout = os.Stdout
	execCommand.Stderr = os.Stderr
	err = execCommand.Run()
	if err != nil {
		log.Printf("Error occurred while executing callback file: %v\n", err)
		return
	}

	log.Printf("Error occurred while executing callback file: %v\n", err)
}
