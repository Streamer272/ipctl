package options

import (
	"github.com/Streamer272/ipctl/handle_error"
	"github.com/Streamer272/ipctl/listener"
)

type Options struct {
	Command  string `json:"command"`
	Interval int    `json:"interval"`
	Current  string `json:"current"`
}

func Default() Options {
	ip, err := listener.GetCurrentIp()
	handle_error.HandleError(err)

	return Options{
		Command:  "echo \"Hello World!\"",
		Interval: 60000,
		Current:  ip,
	}
}
