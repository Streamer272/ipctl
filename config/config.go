package config

import (
	"encoding/json"
	"errors"
	"github.com/Streamer272/ipctl/constants"
	"github.com/Streamer272/ipctl/handle_error"
	"github.com/Streamer272/ipctl/options"
	"io/ioutil"
	"strings"
)

func Get(name string) string {
	content, err := ioutil.ReadFile("/etc/ipctl/ipctl.json")
	handle_error.HandleError(err)

	var opt options.Options
	err = json.Unmarshal(content, &opt)
	handle_error.HandleError(err)

	switch strings.ToLower(name) {
	case "command":
		return opt.Command
	case "interval":
		return opt.GetIntervalString()
	case "current":
		return opt.Current
	default:
		handle_error.HandleError(errors.New("name is not valid"))
		return ""
	}
}

func Set(name string, value string) {
	content, err := ioutil.ReadFile("/etc/ipctl/ipctl.json")
	handle_error.HandleError(err)

	var opt options.Options
	err = json.Unmarshal(content, &opt)
	handle_error.HandleError(err)

	switch strings.ToLower(name) {
	case "command":
		opt.Command = value
		break
	case "interval":
		opt.SetIntervalString(value)
		break
	case "current":
		opt.Current = value
		break
	default:
		handle_error.HandleError(errors.New("name is not valid"))
		return
	}

	updated, err := json.Marshal(opt)
	handle_error.HandleError(err)

	err = ioutil.WriteFile("/etc/ipctl/ipctl.json", []byte(updated), constants.PERMS)
	handle_error.HandleError(err)
}
