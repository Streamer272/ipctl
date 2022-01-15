package systemctl

func Enable() {
	execSystemctl("enable")
}

func Disable() {
	execSystemctl("disable")
}

func Status() {
	execSystemctl("status")
}

func Start() {
	execSystemctl("start")
}

func Stop() {
	execSystemctl("stop")
}

func Restart() {
	execSystemctl("restart")
}

func Reload() {
	execSystemctl("reload")
}
