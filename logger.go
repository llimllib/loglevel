// Copyright 2012 The Oneslang Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"./layout"
	"./level"
	"io"
	"log"
)

type Logger struct {
	level  int
	logger *log.Logger
}

// New creates a new Logger.
func New(out io.Writer) *Logger {
	return &Logger{level: level.Debug, logger: log.New(out, "", log.LstdFlags)}
}

// Sets the output prefix for the logger.
func (l *Logger) setPrefix(funcs int) {
	if l.logger.Flags()&layout.Level != 0 {
		switch funcs {
		case level.Fatal:
			l.logger.SetPrefix("Fatal ")
		case level.Error:
			l.logger.SetPrefix("Error ")
		case level.Warn:
			l.logger.SetPrefix("Warn ")
		case level.Info:
			l.logger.SetPrefix("Info ")
		case level.Debug:
			l.logger.SetPrefix("Debug ")
		case level.Trace:
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
	if funcs <= l.level {
		l.setPrefix(funcs)
		l.logger.Print(v)
	}
}

// Calls Output to printf to the logger.
func (l *Logger) printf(funcs int, format string, v ...interface{}) {
	if funcs <= l.level {
		l.setPrefix(funcs)
		l.logger.Printf(format, v)
	}
}

// Calls Output to println to the logger.
func (l *Logger) println(funcs int, v ...interface{}) {
	if funcs <= l.level {
		l.setPrefix(funcs)
		l.logger.Println(v)
	}
}

// Levels returns the output Levels for the logger.
func (l *Logger) Levels() int {
	return l.level
}

// SetLevels sets the output Levels for the logger.
func (l *Logger) SetLevels(levels int) {
	l.level = levels
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
	l.print(level.Fatal, v)
}

// Calls Output to printf to the logger with the Fatal level.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.printf(level.Fatal, format, v)
}

// Calls Output to println to the logger with the Fatal level.
func (l *Logger) Fatalln(v ...interface{}) {
	l.println(level.Fatal, v)
}

// Calls Output to print to the logger with the Error level.
func (l *Logger) Error(v ...interface{}) {
	l.print(level.Error, v)
}

// Calls Output to printf to the logger with the Error level.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.printf(level.Error, format, v)
}

// Calls Output to println to the logger with the Error level.
func (l *Logger) Errorln(v ...interface{}) {
	l.println(level.Error, v)
}

// Calls Output to print to the logger with the Warn level.
func (l *Logger) Warn(v ...interface{}) {
	l.print(level.Warn, v)
}

// Calls Output to printf to the logger with the Warn level.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.printf(level.Warn, format, v)
}

// Calls Output to println to the logger with the Warn level.
func (l *Logger) Warnln(v ...interface{}) {
	l.println(level.Warn, v)
}

// Calls Output to print to the logger with the Info level.
func (l *Logger) Info(v ...interface{}) {
	l.print(level.Info, v)
}

// Calls Output to printf to the logger with the Info level.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.printf(level.Info, format, v)
}

// Calls Output to println to the logger with the Info level.
func (l *Logger) Infoln(v ...interface{}) {
	l.println(level.Info, v)
}

// Calls Output to print to the logger with the Debug level.
func (l *Logger) Debug(v ...interface{}) {
	l.print(level.Debug, v)
}

// Calls Output to printf to the logger with the Debug level.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.printf(level.Debug, format, v)
}

// Calls Output to println to the logger with the Debug level.
func (l *Logger) Debugln(v ...interface{}) {
	l.println(level.Debug, v)
}

// Calls Output to print to the logger with the Trace level.
func (l *Logger) Trace(v ...interface{}) {
	l.print(level.Trace, v)
}

// Calls Output to printf to the logger with the Trace level.
func (l *Logger) Tracef(format string, v ...interface{}) {
	l.printf(level.Trace, format, v)
}

// Calls Output to println to the logger with the Trace level.
func (l *Logger) Traceln(v ...interface{}) {
	l.println(level.Trace, v)
}
