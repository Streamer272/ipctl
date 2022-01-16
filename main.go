package main

import (
	"fmt"
	"github.com/Streamer272/ipctl/config"
	"github.com/Streamer272/ipctl/constants"
	"github.com/Streamer272/ipctl/handle_error"
	"github.com/Streamer272/ipctl/listener"
	"github.com/Streamer272/ipctl/systemctl"
	"github.com/akamensky/argparse"
	"os"
	"runtime"
	"strconv"
)

func main() {
	if runtime.GOOS != "linux" {
		fmt.Printf("ipctl is only available on linux\n")
		os.Exit(1)
	}

	parser := argparse.NewParser("ipctl", "IP controller\nListen to IP change and change your DNS' records dynamically")

	helpCommand := parser.NewCommand("help", "Display help message")
	versionCommand := parser.NewCommand("version", "Display program version")
	versionFlag := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Display program version", Default: false})

	configCommand := parser.NewCommand("config", "Manage configuration")
	getCommand := configCommand.NewCommand("get", "Show current configuration")
	field := getCommand.String("f", "field", &argparse.Options{Required: false, Help: "Field to get", Default: ""})
	setCommand := configCommand.NewCommand("set", "Update configuration")
	name := setCommand.String("n", "name", &argparse.Options{Required: true, Help: "Name of field to update"})
	value := setCommand.String("V", "value", &argparse.Options{Required: true, Help: "New value"})
	initConfigCommand := configCommand.NewCommand("init", "Initialize configuration")
	rewriteCommand := configCommand.NewCommand("rewrite", "Rewrite configuration to default")
	removeConfigCommand := configCommand.NewCommand("remove", "Remove configuration")

	serviceCommand := parser.NewCommand("service", "Manage service")
	initServiceCommand := serviceCommand.NewCommand("init", "Initialize service")
	enableFlag := initServiceCommand.Flag("e", "enable", &argparse.Options{Required: false, Help: "Enable systemctl service", Default: false})
	removeServiceCommand := serviceCommand.NewCommand("remove", "Remove service")
	enableCommand := serviceCommand.NewCommand("enable", "Enable listening service")
	disableCommand := serviceCommand.NewCommand("disable", "Disable listening service")
	statusCommand := serviceCommand.NewCommand("status", "Status of listening service")
	startCommand := serviceCommand.NewCommand("start", "Start listening service")
	stopCommand := serviceCommand.NewCommand("stop", "Stop listening service")
	restartCommand := serviceCommand.NewCommand("restart", "Restart listening service")
	reloadCommand := serviceCommand.NewCommand("reload", "Reload listening service")
	logsCommand := serviceCommand.NewCommand("logs", "Show service logs")

	listenCommand := parser.NewCommand("listen", "Listen to IP change")

	err := parser.Parse(os.Args)
	if err != nil || helpCommand.Happened() {
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}
	if *versionFlag || versionCommand.Happened() {
		fmt.Printf("%v version %v\n", parser.GetName(), constants.VERSION)
		os.Exit(0)
	}

	if os.Getuid() != 0 {
		fmt.Printf("You need to run %v as root\n", parser.GetName())
		os.Exit(1)
	}

	if configCommand.Happened() {
		if getCommand.Happened() {
			if *field == "" {
				fmt.Printf("%v\n", config.GetAll())
			} else {
				fmt.Printf("%v\n", config.Get(*field))
			}
		}
		if setCommand.Happened() {
			config.Set(*name, *value)
		}
		if initConfigCommand.Happened() {
			config.Init(false)
		}
		if rewriteCommand.Happened() {
			config.Init(true)
		}
		if removeConfigCommand.Happened() {
			config.Remove()
		}
	}

	if listenCommand.Happened() {
		interval, err := strconv.Atoi(config.Get("interval"))
		handle_error.HandleError(err)

		listener.Listen(config.Get("command"), interval)
	}

	if serviceCommand.Happened() {
		if initServiceCommand.Happened() {
			systemctl.Init(*enableFlag)
		}
		if removeServiceCommand.Happened() {
			systemctl.Remove()
		}
		if enableCommand.Happened() {
			systemctl.Enable()
		}
		if disableCommand.Happened() {
			systemctl.Disable()
		}
		if statusCommand.Happened() {
			systemctl.Status()
		}
		if startCommand.Happened() {
			systemctl.Start()
		}
		if stopCommand.Happened() {
			systemctl.Stop()
		}
		if restartCommand.Happened() {
			systemctl.Restart()
		}
		if reloadCommand.Happened() {
			systemctl.Reload()
		}
		if logsCommand.Happened() {
			systemctl.Logs()
		}
	}
}
