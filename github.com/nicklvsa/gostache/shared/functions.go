package shared

import "fmt"

//Log - log something to console base on severity
func Log(severity LogLevel, message interface{}) {
	var severityStr string

	switch severity {
	case Error:
		severityStr = "Error"
	case Warn:
		severityStr = "Warning"
	case Info:
		severityStr = "Info"
	default:
		severityStr = "Info"
	}

	if severityStr != "" && message != "" {
		fmt.Println(fmt.Sprintf("[GOSTACHE] (%s) -> %s", severityStr, message))
	}
}

func 