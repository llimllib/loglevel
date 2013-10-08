// Copyright 2012 The Oneslang Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"io"
	"log"
)

type Logger struct {
	priority int
	prefix   string
	logger   *log.Logger
}

// New creates a new Logger.
func New(out io.Writer) *Logger {
	return &Logger{priority: Pdebug, logger: log.New(out, "", log.LstdFlags)}
}

// Sets the output prefix for the logger.
func (me *Logger) SetPrefix(prefix string) {
	me.prefix = prefix
	me.logger.SetPrefix(prefix)
}

func (me *Logger) setFullPrefix(priority int) {
	if me.logger.Flags()&Lpriority != 0 {
		me.logger.SetPrefix(fmt.Sprintf("%s ", priorityName[priority]) + me.prefix)
	}
}

// Calls Output to print to the logger.
func (me *Logger) print(priority int, v ...interface{}) {
	if priority <= me.priority {
		me.setFullPrefix(priority)
		me.logger.Print(v...)
	}
}

// Calls Output to printf to the logger.
func (me *Logger) printf(priority int, format string, v ...interface{}) {
	if priority <= me.priority {
		me.setFullPrefix(priority)
		me.logger.Printf(format, v...)
	}
}

// Calls Output to println to the logger.
func (me *Logger) println(priority int, v ...interface{}) {
	if priority <= me.priority {
		me.setFullPrefix(priority)
		me.logger.Println(v...)
	}
}

// Priority returns the output priority for the logger.
func (me *Logger) Priority() int {
	return me.priority
}

// SetPriority sets the output priority for the logger.
func (me *Logger) SetPriority(priority int) {
	me.priority = priority
}

// Flags returns the output layouts for the logger.
func (me *Logger) Flags() int {
	return me.logger.Flags()
}

// SetFlags sets the output layouts for the logger.
func (me *Logger) SetFlags(layouts int) {
	me.logger.SetFlags(layouts)
}

// Calls Output to print to the logger with the Fatal level.
func (me *Logger) Fatal(v ...interface{}) {
	me.setFullPrefix(Pfatal)
	me.logger.Fatal(v...)
}

// Calls Output to printf to the logger with the Fatal level.
func (me *Logger) Fatalf(format string, v ...interface{}) {
	me.setFullPrefix(Pfatal)
	me.logger.Fatalf(format, v...)
}

// Calls Output to println to the logger with the Fatal level.
func (me *Logger) Fatalln(v ...interface{}) {
	me.setFullPrefix(Pfatal)
	me.logger.Fatalln(v...)
}

// Calls Output to print to the logger with the Error level.
func (me *Logger) Error(v ...interface{}) {
	me.print(Perror, v...)
}

// Calls Output to printf to the logger with the Error level.
func (me *Logger) Errorf(format string, v ...interface{}) {
	me.printf(Perror, format, v...)
}

// Calls Output to println to the logger with the Error level.
func (me *Logger) Errorln(v ...interface{}) {
	me.println(Perror, v...)
}

// Calls Output to print to the logger with the Warn level.
func (me *Logger) Warn(v ...interface{}) {
	me.print(Pwarn, v...)
}

// Calls Output to printf to the logger with the Warn level.
func (me *Logger) Warnf(format string, v ...interface{}) {
	me.printf(Pwarn, format, v...)
}

// Calls Output to println to the logger with the Warn level.
func (me *Logger) Warnln(v ...interface{}) {
	me.println(Pwarn, v...)
}

// Calls Output to print to the logger with the Info level.
func (me *Logger) Info(v ...interface{}) {
	me.print(Pinfo, v...)
}

// Calls Output to printf to the logger with the Info level.
func (me *Logger) Infof(format string, v ...interface{}) {
	me.printf(Pinfo, format, v...)
}

// Calls Output to println to the logger with the Info level.
func (me *Logger) Infoln(v ...interface{}) {
	me.println(Pinfo, v...)
}

// Calls Output to print to the logger with the Debug level.
func (me *Logger) Debug(v ...interface{}) {
	me.print(Pdebug, v...)
}

// Calls Output to printf to the logger with the Debug level.
func (me *Logger) Debugf(format string, v ...interface{}) {
	me.printf(Pdebug, format, v...)
}

// Calls Output to println to the logger with the Debug level.
func (me *Logger) Debugln(v ...interface{}) {
	me.println(Pdebug, v...)
}

// Calls Output to print to the logger with the Trace level.
func (me *Logger) Trace(v ...interface{}) {
	me.print(Ptrace, v...)
}

// Calls Output to printf to the logger with the Trace level.
func (me *Logger) Tracef(format string, v ...interface{}) {
	me.printf(Ptrace, format, v...)
}

// Calls Output to println to the logger with the Trace level.
func (me *Logger) Traceln(v ...interface{}) {
	me.println(Ptrace, v...)
}
