package main

import (
	"fmt"
	"github.com/Streamer272/ipctl/listener"
	commandParser "github.com/Streamer272/ipctl/parser"
	"github.com/akamensky/argparse"
	"os"
)

const VERSION = "1.0"

func main() {
	parser := argparse.NewParser("ipctl", "IP controller\nListen to IP change and change your DNS' records dynamically")

	interval := parser.Int("i", "interval", &argparse.Options{Required: false, Help: "Request interval", Default: 60000})
	command := parser.String("c", "command", &argparse.Options{Required: false, Help: "IP change command", Default: "echo \"IP changed to $IP!\""})
	version := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Display program version", Default: false})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *version {
		fmt.Printf("%v version %v\n", parser.GetName(), VERSION)
		os.Exit(0)
	}

	if os.Getuid() != 0 {
		fmt.Printf("You need to run %v as root\n", parser.GetName())
		os.Exit(1)
	}

	parsedCommand, args := commandParser.ParseCommand(*command)

	listener.Listen(*interval, parsedCommand, args)
}
