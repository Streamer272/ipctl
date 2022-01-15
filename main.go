package main

import (
	"fmt"
	"github.com/Streamer272/ipctl/config"
	"github.com/Streamer272/ipctl/handle_error"
	"github.com/Streamer272/ipctl/listener"
	"github.com/akamensky/argparse"
	"os"
	"runtime"
	"strconv"
)

const VERSION = "1.0"

func main() {
	if runtime.GOOS != "linux" {
		fmt.Printf("ipctl only available on linux\n")
		os.Exit(1)
	}

	parser := argparse.NewParser("ipctl", "IP controller\nListen to IP change and change your DNS' records dynamically")

	helpCommand := parser.NewCommand("help", "Display help message")
	versionCommand := parser.NewCommand("version", "Display program version")

	initCommand := parser.NewCommand("init", "Initialize ipctl")
	listenCommand := parser.NewCommand("listen", "Listen to IP change")

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

	if initCommand.Happened() {
		config.Init()
	}
	if listenCommand.Happened() {
		interval, err := strconv.Atoi(config.Get("interval"))
		handle_error.HandleError(err)

		listener.Listen(config.Get("command"), interval)
	}
}
