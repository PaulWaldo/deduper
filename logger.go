package main

import (
	"fmt"
	"os"
)

// Logger is responsible for providing feedback to the user
type Logger struct {
	// Whether to show verbose output
	IsVerbose bool
}

// NewLogger creates a new Logger instance
func NewLogger(verbose bool) *Logger {
	return &Logger{
		IsVerbose: verbose,
	}
}

// SetVerbose sets the verbose mode
func (l *Logger) SetVerbose(verbose bool) {
	l.IsVerbose = verbose
}

// Info logs a message regardless of verbosity
func (l *Logger) Info(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

// LogVerbose logs a message only in verbose mode
func (l *Logger) LogVerbose(format string, args ...interface{}) {
	if l.IsVerbose {
		fmt.Printf(format+"\n", args...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
}
