// Copyright 2012 The Oneslang Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package log implements a simple logging package.
package log

import (
	"os"
)

var std *Logger = New(os.Stderr)

// Priority returns the output priority for the standard logger.
func Priority() int {
	return std.Priority()
}

// SetPriority sets the output priority for the standard logger.
func SetPriority(priority int) {
	std.SetPriority(priority)
}

// Layouts returns the output layouts for the standard logger.
func Layouts() int {
	return std.Layouts()
}

// SetLayouts sets the output layouts for the standard logger.
func SetLayouts(layouts int) {
	std.SetLayouts(layouts)
}

// Calls Output to print to the standard logger with the Fatal level.
func Fatal(v ...interface{}) {
	std.Fatal(v)
}

// Calls Output to printf to the standard logger with the Fatal level.
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v)
}

// Calls Output to println to the standard logger with the Fatal level.
func Fatalln(v ...interface{}) {
	std.Fatalln(v)
}

// Calls Output to print to the standard logger with the Error level.
func Error(v ...interface{}) {
	std.Error(v)
}

// Calls Output to printf to the standard logger with the Error level.
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v)
}

// Calls Output to println to the standard logger with the Error level.
func Errorln(v ...interface{}) {
	std.Errorln(v)
}

// Calls Output to print to the standard logger with the Warn level.
func Warn(v ...interface{}) {
	std.Warn(v)
}

// Calls Output to printf to the standard logger with the Warn level.
func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v)
}

// Calls Output to println to the standard logger with the Warn level.
func Warnln(v ...interface{}) {
	std.Warnln(v)
}

// Calls Output to print to the standard logger with the Info level.
func Info(v ...interface{}) {
	std.Info(v)
}

// Calls Output to printf to the standard logger with the Info level.
func Infof(format string, v ...interface{}) {
	std.Infof(format, v)
}

// Calls Output to println to the standard logger with the Info level.
func Infoln(v ...interface{}) {
	std.Infoln(v)
}

// Calls Output to print to the standard logger with the Debug level.
func Debug(v ...interface{}) {
	std.Debug(v)
}

// Calls Output to printf to the standard logger with the Debug level.
func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v)
}

// Calls Output to println to the standard logger with the Debug level.
func Debugln(v ...interface{}) {
	std.Debugln(v)
}

// Calls Output to print to the standard logger with the Trace level.
func Trace(v ...interface{}) {
	std.Trace(v)
}

// Calls Output to printf to the standard logger with the Trace level.
func Tracef(format string, v ...interface{}) {
	std.Tracef(format, v)
}

// Calls Output to println to the standard logger with the Trace level.
func Traceln(v ...interface{}) {
	std.Traceln(v)
}
