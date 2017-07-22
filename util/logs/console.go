package logs

import (
	"log"
	"os"
)

func NewConsoleLogger(level LogLevel) *ConsoleLogger {
	return &ConsoleLogger{
		level:  level,
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

type ConsoleLogger struct {
	level LogLevel
	*log.Logger
}

func (cl *ConsoleLogger) D(format string, args ...interface{}) {
	cl.output(3, DEBUG, format, args...)
}

func (cl *ConsoleLogger) I(format string, args ...interface{}) {
	cl.output(3, INFO, format, args...)
}

func (cl *ConsoleLogger) W(format string, args ...interface{}) {
	cl.output(3, WARNING, format, args...)
}

func (cl *ConsoleLogger) E(format string, args ...interface{}) {
	cl.output(3, ERROR, format, args...)
}

func (cl *ConsoleLogger) F(format string, args ...interface{}) {
	cl.output(3, FATAL, format, args...)
}

func (cl *ConsoleLogger) SetTraceLevel(level LogLevel) {
	cl.level = level
}

func (cl *ConsoleLogger) Flush() error {
	return nil
}

func (cl *ConsoleLogger) output(depth int, level LogLevel, format string, args ...interface{}) {
	if level < cl.level {
		return
	}
	cl.Output(depth, getString(level, format, args...))
}
