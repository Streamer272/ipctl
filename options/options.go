package options

import (
	"github.com/Streamer272/ipctl/handle_error"
	"github.com/Streamer272/ipctl/listener"
	"strconv"
)

type Options struct {
	Command  string `json:"command"`
	interval string `json:"interval"`
	Current  string `json:"current"`
}

func (this *Options) GetInterval() int {
	intValue, err := strconv.Atoi(this.interval)
	handle_error.HandleError(err)
	return intValue
}

func (this *Options) GetIntervalString() string {
	return this.interval
}

func (this *Options) SetInterval(value int) {
	this.interval = strconv.Itoa(value)
}

func (this *Options) SetIntervalString(value string) {
	this.interval = value
}

func Default() Options {
	ip, err := listener.GetCurrentIp()
	handle_error.HandleError(err)

	return Options{
		Command:  "echo \"Hello World!\"",
		interval: "60000",
		Current:  ip,
	}
}
