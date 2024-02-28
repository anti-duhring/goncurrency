package logger

import "fmt"

func logEntry(level string, msg string) {

	logEntry := fmt.Sprintf("[%s] %s", level, msg)

	fmt.Println(logEntry)
}

func Error(msg string, err error) {
	logEntry("ERROR", fmt.Sprintf("%s: %s", msg, err))
}

func Debug(msg string) {
	logEntry("DEBUG", msg)
}

