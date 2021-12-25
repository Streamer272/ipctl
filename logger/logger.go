package logger

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

type Logger struct {
	disabled bool
}

func (log *Logger) Log(logLevel string, message string) {
	if logLevelToInt(logLevel) == 3 /* none */ || log.disabled {
		return
	}

	currentTime := time.Now()
	output := os.Stdout
	if logLevelToInt(logLevel) == 2 /* error */ {
		output = os.Stderr
	}
	_, err := fmt.Fprintf(output, "[%v] %v: %v\n", getColorFuncByLogLevel(logLevelToInt(logLevel))(logLevel), currentTime.Format("15:04:05"), message)
	if err != nil {
		fmt.Printf("Couldn't log data, err = %v\n", err)
	}
}

func (log *Logger) Disable() {
	log.disabled = true
}

func (log *Logger) Enable() {
	log.disabled = false
}

func NewLogger() Logger {
	return Logger{
		disabled: false,
	}
}

func logLevelToInt(logLevel string) int {
	switch logLevel {
	case "INFO":
		return 0
	case "WARN":
		return 1
	case "ERROR":
		return 2
	default:
		return 3
	}
}

func getColorFuncByLogLevel(logLevel int) func(a ...interface{}) string {
	switch logLevel {
	case 0:
		return color.New(color.FgGreen).SprintFunc()
	case 1:
		return color.New(color.FgYellow).SprintFunc()
	case 2:
		return color.New(color.FgRed).SprintFunc()
	default:
		return color.New(color.FgWhite).SprintFunc()
	}
}
