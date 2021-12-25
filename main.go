package main

import (
	"fmt"
	"github.com/Streamer272/ipcl/listener"
	"github.com/akamensky/argparse"
	"os"
	"strings"
)

func main() {
	parser := argparse.NewParser("ipcl", "IP change listener")

	interval := parser.Int("i", "interval", &argparse.Options{Required: false, Help: "Request interval", Default: 5000})
	callback := parser.String("c", "callback", &argparse.Options{Required: false, Help: "IP change callback", Default: "echo \"IP changed to $IP!\""})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	command := strings.Split(*callback, " ")[0]
	var args []string
	currentArg := ""
	for _, arg := range strings.Split(*callback, " ")[1:] {
		if strings.Count(arg, "\"")%2 == 1 && currentArg != "" {
			currentArg = arg + " "
		} else {
			args = append(args, currentArg+arg)
			currentArg = ""
		}
	}

	listener.Listen(*interval, command, args)
}
