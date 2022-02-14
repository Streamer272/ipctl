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
	"github.com/Streamer272/ipctl/systemctl"
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
			fmt.Printf("%v\n", ip)
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

	app.Command("service", "Manage ipctl service", func(cmd *cli.Cmd) {
		cmdToDo := cmd.StringOpt("x execute", "", "Command to execute with systemctl")
		noOutput := cmd.BoolOpt("n no-output", false, "Not to display output")

		cmd.Command("init", "Initialize ipctl service", func(cmd *cli.Cmd) {
			noEnable := cmd.BoolOpt("n no-enable", false, "Not to enable service")

			cmd.Action = func() {
				systemctl.Init(!*noEnable)
			}
		})

		cmd.Action = func() {
			if strings.HasPrefix(*cmdToDo, "log") {
				systemctl.Logs()
			} else if *cmdToDo != "" {
				systemctl.Exec(*cmdToDo, !*noOutput)
			} else {
				fmt.Printf("Edit %v to change service configuration\n", "/lib/systemd/system/ipctl.service")
			}
		}
	})

	app.Run(os.Args)
}
