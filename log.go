// Loglevel is a simple levelled logging wrapper around the go "log" module
package loglevel

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var std *Logger = New(os.Stderr, "", LstdFlags, Pdebug)

func SetOutput(out io.Writer) {
	std = New(out, std.Prefix(), std.Flags(), std.priority)
}

// Priority returns the output priority for the standard logger.
func Priority() int {
	return std.Priority()
}

// SetPriority sets the output priority for the standard logger.
func SetPriority(priority int) {
	std.SetPriority(priority)
}

// SetPriorityString sets the output priority by the name of a debug level
func SetPriorityString(s string) error {
	s = strings.ToUpper(s)
	for i, name := range priorityName {
		if name == s {
			SetPriority(i)
			return nil
		}
	}
	return fmt.Errorf("Unable to find priority %s", s)
}

// Flags returns the output layouts for the standard logger.
func Flags() int {
	return std.Flags()
}

// SetFlags sets the output layouts for the standard logger.
func SetFlags(flags int) {
	std.SetFlags(flags)
}

// SetPrefix sets the logger prefix
func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

// Calls Output to print to the standard logger with the Fatal level.
func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

// Calls Output to printf to the standard logger with the Fatal level.
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

// Calls Output to println to the standard logger with the Fatal level.
func Fatalln(v ...interface{}) {
	std.Fatalln(v...)
}

// Calls Output to print to the standard logger with the Panic level.
func Panic(v ...interface{}) {
	std.Panic(v...)
}

// Calls Output to printf to the standard logger with the Panic level.
func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

// Calls Output to println to the standard logger with the Panic level.
func Panicln(v ...interface{}) {
	std.Panicln(v...)
}

// Calls Output to print to the standard logger with the Error level.
func Error(v ...interface{}) {
	std.Error(v...)
}

// Calls Output to printf to the standard logger with the Error level.
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

// Calls Output to println to the standard logger with the Error level.
func Errorln(v ...interface{}) {
	std.Errorln(v...)
}

// Calls Output to print to the standard logger with the Warn level.
func Warn(v ...interface{}) {
	std.Warn(v...)
}

// Calls Output to printf to the standard logger with the Warn level.
func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v...)
}

// Calls Output to println to the standard logger with the Warn level.
func Warnln(v ...interface{}) {
	std.Warnln(v...)
}

// Calls Output to print to the standard logger with the Info level.
func Info(v ...interface{}) {
	std.Info(v...)
}

// Calls Output to printf to the standard logger with the Info level.
func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

// Calls Output to println to the standard logger with the Info level.
func Infoln(v ...interface{}) {
	std.Infoln(v...)
}

// Calls Output to print to the standard logger with the Debug level.
func Debug(v ...interface{}) {
	std.Debug(v...)
}

// Calls Output to printf to the standard logger with the Debug level.
func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

// Calls Output to println to the standard logger with the Debug level.
func Debugln(v ...interface{}) {
	std.Debugln(v...)
}

// Calls Output to print to the standard logger with the Trace level.
func Trace(v ...interface{}) {
	std.Trace(v...)
}

// Calls Output to printf to the standard logger with the Trace level.
func Tracef(format string, v ...interface{}) {
	std.Tracef(format, v...)
}

// Calls Output to println to the standard logger with the Trace level.
func Traceln(v ...interface{}) {
	std.Traceln(v...)
}
