package main

import (
	"fmt"
	"github.com/Streamer272/ipctl/config"
	"github.com/akamensky/argparse"
	"os"
)

const VERSION = "1.0"

func main() {
	parser := argparse.NewParser("ipctl", "IP controller\nListen to IP change and change your DNS' records dynamically")

	helpCommand := parser.NewCommand("help", "Display help message")
	versionCommand := parser.NewCommand("version", "Display program version")

	versionFlag := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Display program version", Default: false})

	err := parser.Parse(os.Args)
	if err != nil || helpCommand.Happened() {
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}

	if *versionFlag || versionCommand.Happened() {
		fmt.Printf("%v versionFlag %v\n", parser.GetName(), VERSION)
		os.Exit(0)
	}

	if os.Getuid() != 0 {
		fmt.Printf("You need to run %v as root\n", parser.GetName())
		os.Exit(1)
	}

	config.Get("test")

	//parsedCommand, args := commandParser.ParseCommand(*command)
	//listener.Listen(*interval, parsedCommand, args)
}
