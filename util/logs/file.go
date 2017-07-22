package logs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func NewFileLogger(level LogLevel) *FileLogger {
	return &FileLogger{
		level:     level,
		Logger:    log.New(os.Stdout, "", log.LstdFlags),
		//bufWriter: bufio.NewWriter(os.Stdout.),
	}
}

type FileLogger struct {
	level             LogLevel
	bufWriter         bufio.Writer
	fileNameFormat    string
	errFileNameFormat string
	currentFile       string
	currentErrFile    string
	*log.Logger
}

func (fl *FileLogger) D(format string, args ...interface{}) {
	fl.output(3, DEBUG, format, args...)
}

func (fl *FileLogger) I(format string, args ...interface{}) {
	fl.output(3, INFO, format, args...)
}

func (fl *FileLogger) W(format string, args ...interface{}) {
	fl.output(3, WARNING, format, args...)
}

func (fl *FileLogger) E(format string, args ...interface{}) {
	fl.output(3, ERROR, format, args...)
}

func (fl *FileLogger) F(format string, args ...interface{}) {
	fl.output(3, FATAL, format, args...)
}

func (fl *FileLogger) SetTraceLevel(level LogLevel) {
	fl.level = level
}

func (fl *FileLogger) Flush() error {
	return nil
}

func (fl *FileLogger) output(depth int, level LogLevel, format string, args ...interface{}) {
	if level < fl.level {
		return
	}
	var name = time.Now().Format("2006-01-02")
	if len(fl.fileNameFormat) > 0 {
		name = fmt.Sprintf(fl.fileNameFormat, name)
	}
	if name != fl.currentFile {

	}
	fl.Output(depth, getString(level, format, args...))
}

func (fl *FileLogger) open() {

}
