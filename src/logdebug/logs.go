package logs

import (
	"github.com/cxr29/log"
	"io"
	"os"
)

const (
	LevelDebug = (iota + 1) * 100
	LevelInfo
	LevelNotice
	LevelWarning
	LevelError
	LevelCritical
	LevelPanic
	LevelFatal
)

var DEBUG = log.Debug
var INFO = log.Info
var ERROR = log.Error
var WARN = log.Warning

func InitNewLogFile(logfile string) *os.File {
	logFile, err := os.Create(logfile)

	if err != nil {
		panic(err)
	}
	writers := []io.Writer{
		logFile,
		os.Stdout,
	}

	fileAndStdout := io.MultiWriter(writers...)
	log.SetOutput(fileAndStdout)
	return logFile
}

func SetLogLevel(level int) {
	log.SetLevel(level)
}
