package config

import (
	"encoding/json"
	"github.com/Streamer272/ipctl/handle_error"
	"github.com/Streamer272/ipctl/options"
	"io/ioutil"
)

func Get(name string) interface{} {
	content, err := ioutil.ReadFile("/etc/ipctl/ipctl.json")
	handle_error.HandleError(err)

	var opt options.Options
	err = json.Unmarshal(content, &opt)
	handle_error.HandleError(err)

	return 0
}
