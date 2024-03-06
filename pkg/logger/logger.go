package logger

import (
	"log"
	"os"
)

var (
	// InfoLogger is used for logging informational messages.
	InfoLogger *log.Logger

	// ErrorLogger is used for logging error messages.
	ErrorLogger *log.Logger
)

func init() {
	// Initialize loggers with different prefixes and output destinations
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
