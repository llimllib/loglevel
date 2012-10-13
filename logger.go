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
func (l *Logger) setPrefix(funcs int) {
	if l.logger.Flags()&Lpriority != 0 {
		switch funcs {
		case Pfatal:
			l.logger.SetPrefix("Fatal ")
		case Perror:
			l.logger.SetPrefix("Error ")
		case Pwarn:
			l.logger.SetPrefix("Warn ")
		case Pinfo:
			l.logger.SetPrefix("Info ")
		case Pdebug:
			l.logger.SetPrefix("Debug ")
		case Ptrace:
			l.logger.SetPrefix("Trace ")
		default:
			l.logger.SetPrefix("")
		}
	} else {
		l.logger.SetPrefix("")
	}
}

// Calls Output to print to the logger.
func (l *Logger) print(funcs int, v ...interface{}) {
	if funcs <= l.priority {
		l.setPrefix(funcs)
		l.logger.Print(v)
	}
}

// Calls Output to printf to the logger.
func (l *Logger) printf(funcs int, format string, v ...interface{}) {
	if funcs <= l.priority {
		l.setPrefix(funcs)
		l.logger.Printf(format, v)
	}
}

// Calls Output to println to the logger.
func (l *Logger) println(funcs int, v ...interface{}) {
	if funcs <= l.priority {
		l.setPrefix(funcs)
		l.logger.Println(v)
	}
}

// Priority returns the output priority for the logger.
func (l *Logger) Priority() int {
	return l.priority
}

// SetPriority sets the output priority for the logger.
func (l *Logger) SetPriority(priority int) {
	l.priority = priority
}

// Layouts returns the output layouts for the logger.
func (l *Logger) Layouts() int {
	return l.logger.Flags()
}

// SetLayouts sets the output layouts for the logger.
func (l *Logger) SetLayouts(layouts int) {
	l.logger.SetFlags(layouts)
}

// Calls Output to print to the logger with the Fatal level.
func (l *Logger) Fatal(v ...interface{}) {
	l.print(Pfatal, v)
}

// Calls Output to printf to the logger with the Fatal level.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.printf(Pfatal, format, v)
}

// Calls Output to println to the logger with the Fatal level.
func (l *Logger) Fatalln(v ...interface{}) {
	l.println(Pfatal, v)
}

// Calls Output to print to the logger with the Error level.
func (l *Logger) Error(v ...interface{}) {
	l.print(Perror, v)
}

// Calls Output to printf to the logger with the Error level.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.printf(Perror, format, v)
}

// Calls Output to println to the logger with the Error level.
func (l *Logger) Errorln(v ...interface{}) {
	l.println(Perror, v)
}

// Calls Output to print to the logger with the Warn level.
func (l *Logger) Warn(v ...interface{}) {
	l.print(Pwarn, v)
}

// Calls Output to printf to the logger with the Warn level.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.printf(Pwarn, format, v)
}

// Calls Output to println to the logger with the Warn level.
func (l *Logger) Warnln(v ...interface{}) {
	l.println(Pwarn, v)
}

// Calls Output to print to the logger with the Info level.
func (l *Logger) Info(v ...interface{}) {
	l.print(Pinfo, v)
}

// Calls Output to printf to the logger with the Info level.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.printf(Pinfo, format, v)
}

// Calls Output to println to the logger with the Info level.
func (l *Logger) Infoln(v ...interface{}) {
	l.println(Pinfo, v)
}

// Calls Output to print to the logger with the Debug level.
func (l *Logger) Debug(v ...interface{}) {
	l.print(Pdebug, v)
}

// Calls Output to printf to the logger with the Debug level.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.printf(Pdebug, format, v)
}

// Calls Output to println to the logger with the Debug level.
func (l *Logger) Debugln(v ...interface{}) {
	l.println(Pdebug, v)
}

// Calls Output to print to the logger with the Trace level.
func (l *Logger) Trace(v ...interface{}) {
	l.print(Ptrace, v)
}

// Calls Output to printf to the logger with the Trace level.
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.printf(Ptrace, format, v)
}

// Calls Output to println to the logger with the Trace level.
func (l *Logger) Traceln(v ...interface{}) {
	l.println(Ptrace, v)
}
