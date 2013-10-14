// Package loglevel implements a simple levelled logging wrapper around the go
// "log" module
//
// The goal of this project is to be the simplest well-tested levelled logging
// wrapper around go's log module. It retains as much of the same API as possible.
package loglevel

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var std = New(os.Stderr, "", LstdFlags, Pdebug)

// SetOutput sets the output destination for the standard logger
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

// Fatal prints the message it's given and quits the program
func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

// Fatalf prints the message it's given and quits the program
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

// Fatalln prints the message it's given and quits the program
func Fatalln(v ...interface{}) {
	std.Fatalln(v...)
}

// Panic prints the message it's given and panic()s the program
func Panic(v ...interface{}) {
	std.Panic(v...)
}

// Panicf prints the message it's given and panic()s the program
func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

// Panicln prints the message it's given and panic()s the program
func Panicln(v ...interface{}) {
	std.Panicln(v...)
}

// Error prints to the standard logger with the Error level.
func Error(v ...interface{}) {
	std.Error(v...)
}

// Errorf prints to the standard logger with the Error level.
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

// Errorln prints to the standard logger with the Error level.
func Errorln(v ...interface{}) {
	std.Errorln(v...)
}

// Warn prints to the standard logger with the Warn level.
func Warn(v ...interface{}) {
	std.Warn(v...)
}

// Warnf prints to the standard logger with the Warn level.
func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v...)
}

// Warnln prints to the standard logger with the Warn level.
func Warnln(v ...interface{}) {
	std.Warnln(v...)
}

// Info prints to the standard logger with the Info level.
func Info(v ...interface{}) {
	std.Info(v...)
}

// Infof prints to the standard logger with the Info level.
func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

// Infoln prints to the standard logger with the Info level.
func Infoln(v ...interface{}) {
	std.Infoln(v...)
}

// Debug prints to the standard logger with the Debug level.
func Debug(v ...interface{}) {
	std.Debug(v...)
}

// Debugf prints to the standard logger with the Debug level.
func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

// Debugln prints to the standard logger with the Debug level.
func Debugln(v ...interface{}) {
	std.Debugln(v...)
}

// Trace prints to the standard logger with the Trace level.
func Trace(v ...interface{}) {
	std.Trace(v...)
}

// Tracef prints to the standard logger with the Trace level.
func Tracef(format string, v ...interface{}) {
	std.Tracef(format, v...)
}

// Traceln prints to the standard logger with the Trace level.
func Traceln(v ...interface{}) {
	std.Traceln(v...)
}
