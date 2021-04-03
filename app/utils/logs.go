package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

const (
	logFileTrace = "trace.log"
	logFileInfo  = "info.log"
	logFileWarn  = "warn.log"
	logFileError = "error.log"
	logFileGral  = "logs.log"
)

func init() {

	logTrace, err := os.OpenFile(logFileTrace, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileTrace, err))
	}
	logInfo, err := os.OpenFile(logFileInfo, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileInfo, err))
	}
	logWarn, err := os.OpenFile(logFileWarn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileWarn, err))
	}
	logError, err := os.OpenFile(logFileError, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileError, err))
	}
	logGral, err := os.OpenFile(logFileGral, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileGral, err))
	}

	Trace = log.New(io.MultiWriter(logTrace, logGral, os.Stdout),
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(io.MultiWriter(logInfo, logGral, os.Stdout),
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(io.MultiWriter(logWarn, logGral, os.Stdout),
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(logError, logGral, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
