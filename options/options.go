package options

import (
	"github.com/Streamer272/ipctl/handle_error"
	"strconv"
)

type Options struct {
	Command  string `json:"command"`
	Interval string `json:"interval"`
	Current  string `json:"current"`
}

func (this *Options) GetInterval() int {
	intValue, err := strconv.Atoi(this.Interval)
	handle_error.HandleError(err)
	return intValue
}

func (this *Options) GetIntervalString() string {
	return this.Interval
}

func (this *Options) SetInterval(value int) {
	this.Interval = strconv.Itoa(value)
}

func (this *Options) SetIntervalString(value string) {
	this.Interval = value
}

func Default() Options {
	return Options{
		Command:  "echo \"Hello World!\"",
		Interval: "60000",
		Current:  "",
	}
}
