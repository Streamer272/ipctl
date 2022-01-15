package config

import (
	"encoding/json"
	"errors"
	"github.com/Streamer272/ipctl/handle_error"
	"github.com/Streamer272/ipctl/options"
	"io/ioutil"
	"strconv"
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
		return strconv.Itoa(opt.Interval)
	case "current":
		return opt.Current
	default:
		handle_error.HandleError(errors.New("name is not valid"))
		return ""
	}
}

func Set(name string, value interface{}) {
	// TODO
}
