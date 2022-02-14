package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/Streamer272/cool/check"
	"github.com/Streamer272/ipctl/config"
	"github.com/Streamer272/ipctl/constants"
	"github.com/Streamer272/ipctl/listener"
	cli "github.com/jawher/mow.cli"
)

func main() {
	if runtime.GOOS != "linux" {
		fmt.Printf("ipctl is only available on linux\n")
		os.Exit(1)
	}

	app := cli.App("ipctl", "IP controller - Listen for IP change and change your DNS' records dynamically")
	app.Spec = ""
	app.Version("v version", fmt.Sprintf("ipctl version %v", constants.VERSION))

	app.Command("version", "Show the version and exit", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			app.PrintVersion()
		}
	})
	app.Command("help", "Print help and exit", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			app.PrintLongHelp()
		}
	})

	app.Command("ip", "Show current IP", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			ip, err := listener.GetCurrentIp()
			check.Check(err)
			fmt.Printf("%v", ip)
		}
	})

	app.Command("listen", "Listen for IP change", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			interval, err := strconv.Atoi(config.Get("interval"))
			check.Check(err)
			listener.Listen(config.Get("command"), interval)
		}
	})

	app.Command("update", "Update DNS' records", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			listener.Update(config.Get("command"))
		}
	})

	app.Command("config", "Manage ipctl config", func(cmd *cli.Cmd) {
		cmd.Command("init", "Initialize ipctl config", func(cmd *cli.Cmd) {
			cmd.Action = func() {
				config.Init()
			}
		})

		cmd.Action = func() {
			fmt.Printf("Edit %v to change configuration\n", strings.Join(config.GetConfigFiles(), ", "))
		}
	})

	app.Run(os.Args)

	/*
		parser := argparse.NewParser("ipctl", "IP controller\nListen to IP change and change your DNS' records dynamically")

		helpCommand := parser.NewCommand("help", "Display help message")
		versionCommand := parser.NewCommand("version", "Display program version")
		versionFlag := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "Display program version", Default: false})

		ipCommand := parser.NewCommand("ip", "Get current IP")

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
		updateCommand := parser.NewCommand("update", "Update DNS IP address")

		err := parser.Parse(os.Args)
		if err != nil || helpCommand.Happened() {
			fmt.Print(parser.Usage(err))
			os.Exit(0)
		}
		if *versionFlag || versionCommand.Happened() {
			fmt.Printf("%v version %v\n", parser.GetName(), constants.VERSION)
			os.Exit(0)
		}

		if ipCommand.Happened() {
			ip, err := listener.GetCurrentIp()
			if err != nil {
				fmt.Printf("Couldn't get IP, error: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("%v", ip)
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

		if listenCommand.Happened() {
			interval, err := strconv.Atoi(config.Get("interval"))
			handle_error.HandleError(err)

			listener.Listen(config.Get("command"), interval)
		}
		if updateCommand.Happened() {
			listener.Update(config.Get("command"))
		}
	*/
}
