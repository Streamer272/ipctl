package parser

import "strings"

func ParseCommand(command string) (string, []string) {
	var args []string
	currentArg := ""
	for _, arg := range strings.Split(command, " ")[1:] {
		if strings.Count(arg, "\"")%2 == 1 && currentArg != "" {
			currentArg = arg + " "
		} else {
			args = append(args, currentArg+arg)
			currentArg = ""
		}
	}

	return strings.Split(command, " ")[0], args
}
