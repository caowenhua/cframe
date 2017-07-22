package logs

import "fmt"

type LogLevel int

const (
	NONE    LogLevel = 1
	DEBUG            = 2
	INFO             = 3
	WARNING          = 4
	ERROR            = 5
	FATAL            = 6
)

func Rediect(outPath, errPath string) {

}

//var defaultLog =

type ILog interface {
	D(format string, args ...interface{})
	I(format string, args ...interface{})
	W(format string, args ...interface{})
	E(format string, args ...interface{})
	F(format string, args ...interface{})
	SetTraceLevel(level LogLevel)
	Flush() error
}

func getString(level LogLevel, format string, args ...interface{}) string {
	var str string
	switch level {
	case NONE:
		str = fmt.Sprintf("[N] %s", fmt.Sprintf(format, args...)) + "\n"
	case DEBUG:
		str = fmt.Sprintf("[D] %s", fmt.Sprintf(format, args...)) + "\n"
	case INFO:
		str = fmt.Sprintf("[I] %s", fmt.Sprintf(format, args...)) + "\n"
	case WARNING:
		str = fmt.Sprintf("[W] %s", fmt.Sprintf(format, args...)) + "\n"
	case ERROR:
		str = fmt.Sprintf("[E] %s", fmt.Sprintf(format, args...)) + "\n"
	case FATAL:
		str = fmt.Sprintf("[F] %s", fmt.Sprintf(format, args...)) + "\n"
	}
	return str
}
