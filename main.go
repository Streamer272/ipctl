package main

import (
	"fmt"
	"github.com/Streamer272/ipctl/listener"
	"github.com/akamensky/argparse"
	"os"
	"strings"
)

const VERSION = "1.0"

func main() {
	parser := argparse.NewParser("ipctl", "IP controller\nListen to IP change and change your DNS' records dynamically")

	interval := parser.Int("i", "interval", &argparse.Options{Required: false, Help: "Request interval", Default: 60000})
	callback := parser.String("c", "callback", &argparse.Options{Required: false, Help: "IP change callback", Default: "echo \"IP changed to $IP!\""})
	version := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Display program version", Default: false})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *version {
		fmt.Printf("ipctl version %v\n", VERSION)
		os.Exit(0)
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
