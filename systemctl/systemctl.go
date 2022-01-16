package systemctl

func Enable() {
	execSystemctl("enable", true)
}

func Disable() {
	execSystemctl("disable", true)
}

func Status() {
	execSystemctl("status", true)
}

func Start() {
	execSystemctl("start", true)
}

func Stop() {
	execSystemctl("stop", true)
}

func Restart() {
	execSystemctl("restart", true)
}

func Reload() {
	execSystemctl("reload", true)
}
