// Copyright 2012 The Oneslang Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"io"
	"log"
)

type Logger struct {
	priority int
	logger   *log.Logger
}

// New creates a new Logger.
func New(out io.Writer) *Logger {
	return &Logger{priority: Pdebug, logger: log.New(out, "", log.LstdFlags)}
}

// Sets the output prefix for the logger.
func (me *Logger) setPrefix(funcs int) {
	if me.logger.Flags()&Lpriority != 0 {
		switch funcs {
		case Pfatal:
			me.logger.SetPrefix("Fatal ")
		case Perror:
			me.logger.SetPrefix("Error ")
		case Pwarn:
			me.logger.SetPrefix("Warn ")
		case Pinfo:
			me.logger.SetPrefix("Info ")
		case Pdebug:
			me.logger.SetPrefix("Debug ")
		case Ptrace:
			me.logger.SetPrefix("Trace ")
		default:
			me.logger.SetPrefix("")
		}
	} else {
		me.logger.SetPrefix("")
	}
}

// Calls Output to print to the logger.
func (me *Logger) print(funcs int, v ...interface{}) {
	if funcs <= me.priority {
		me.setPrefix(funcs)
		me.logger.Print(v...)
	}
}

// Calls Output to printf to the logger.
func (me *Logger) printf(funcs int, format string, v ...interface{}) {
	if funcs <= me.priority {
		me.setPrefix(funcs)
		me.logger.Printf(format, v...)
	}
}

// Calls Output to println to the logger.
func (me *Logger) println(funcs int, v ...interface{}) {
	if funcs <= me.priority {
		me.setPrefix(funcs)
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

// Layouts returns the output layouts for the logger.
func (me *Logger) Layouts() int {
	return me.logger.Flags()
}

// SetLayouts sets the output layouts for the logger.
func (me *Logger) SetLayouts(layouts int) {
	me.logger.SetFlags(layouts)
}

// Calls Output to print to the logger with the Fatal level.
func (me *Logger) Fatal(v ...interface{}) {
	me.print(Pfatal, v...)
}

// Calls Output to printf to the logger with the Fatal level.
func (me *Logger) Fatalf(format string, v ...interface{}) {
	me.printf(Pfatal, format, v...)
}

// Calls Output to println to the logger with the Fatal level.
func (me *Logger) Fatalln(v ...interface{}) {
	me.println(Pfatal, v...)
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
