package loglevel

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// Logger defines our wrapper around the system logger
type Logger struct {
	priority int
	prefix   string
	logger   *log.Logger
}

// New creates a new Logger.
func New(out io.Writer, prefix string, flag int, priority int) *Logger {
	return &Logger{
		priority: priority,
		prefix:   prefix,
		logger:   log.New(out, prefix, flag),
	}
}

// SetPrefix sets the output prefix for the logger.
func (me *Logger) SetPrefix(prefix string) {
	me.prefix = prefix
	me.logger.SetPrefix(prefix)
}

// Prefix returns the current logger prefix
func (me *Logger) Prefix() string {
	return me.prefix
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

// SetPriorityString sets the output priority by the name of a debug level
func (me *Logger) SetPriorityString(s string) error {
	s = strings.ToUpper(s)
	for i, name := range priorityName {
		if name == s {
			me.SetPriority(i)
			return nil
		}
	}
	return fmt.Errorf("Unable to find priority %s", s)
}

// Flags returns the output layouts for the logger.
func (me *Logger) Flags() int {
	return me.logger.Flags()
}

// SetFlags sets the output layouts for the logger.
func (me *Logger) SetFlags(layouts int) {
	me.logger.SetFlags(layouts)
}

// Fatal prints the message it's given and quits the program
func (me *Logger) Fatal(v ...interface{}) {
	me.setFullPrefix(Pfatal)
	me.logger.Fatal(v...)
}

// Fatalf prints the message it's given and quits the program
func (me *Logger) Fatalf(format string, v ...interface{}) {
	me.setFullPrefix(Pfatal)
	me.logger.Fatalf(format, v...)
}

// Fatalln prints the message it's given and quits the program
func (me *Logger) Fatalln(v ...interface{}) {
	me.setFullPrefix(Pfatal)
	me.logger.Fatalln(v...)
}

// Panic prints the message it's given and panic()s the program
func (me *Logger) Panic(v ...interface{}) {
	me.setFullPrefix(Pfatal)
	me.logger.Panic(v...)
}

// Panicf prints the message it's given and panic()s the program
func (me *Logger) Panicf(format string, v ...interface{}) {
	me.setFullPrefix(Pfatal)
	me.logger.Panicf(format, v...)
}

// Panicln prints the message it's given and panic()s the program
func (me *Logger) Panicln(v ...interface{}) {
	me.setFullPrefix(Pfatal)
	me.logger.Panicln(v...)
}

// Error prints to the standard logger with the Error level.
func (me *Logger) Error(v ...interface{}) {
	me.print(Perror, v...)
}

// Errorf prints to the standard logger with the Error level.
func (me *Logger) Errorf(format string, v ...interface{}) {
	me.printf(Perror, format, v...)
}

// Errorln prints to the standard logger with the Error level.
func (me *Logger) Errorln(v ...interface{}) {
	me.println(Perror, v...)
}

// Warn prints to the standard logger with the Warn level.
func (me *Logger) Warn(v ...interface{}) {
	me.print(Pwarn, v...)
}

// Warnf prints to the standard logger with the Warn level.
func (me *Logger) Warnf(format string, v ...interface{}) {
	me.printf(Pwarn, format, v...)
}

// Warnln prints to the standard logger with the Warn level.
func (me *Logger) Warnln(v ...interface{}) {
	me.println(Pwarn, v...)
}

// Info prints to the standard logger with the Info level.
func (me *Logger) Info(v ...interface{}) {
	me.print(Pinfo, v...)
}

// Infof prints to the standard logger with the Info level.
func (me *Logger) Infof(format string, v ...interface{}) {
	me.printf(Pinfo, format, v...)
}

// Infoln prints to the standard logger with the Info level.
func (me *Logger) Infoln(v ...interface{}) {
	me.println(Pinfo, v...)
}

// Debug prints to the standard logger with the Debug level.
func (me *Logger) Debug(v ...interface{}) {
	me.print(Pdebug, v...)
}

// Debugf prints to the standard logger with the Debug level.
func (me *Logger) Debugf(format string, v ...interface{}) {
	me.printf(Pdebug, format, v...)
}

// Debugln prints to the standard logger with the Debug level.
func (me *Logger) Debugln(v ...interface{}) {
	me.println(Pdebug, v...)
}

// Trace prints to the standard logger with the Trace level.
func (me *Logger) Trace(v ...interface{}) {
	me.print(Ptrace, v...)
}

// Tracef prints to the standard logger with the Trace level.
func (me *Logger) Tracef(format string, v ...interface{}) {
	me.printf(Ptrace, format, v...)
}

// Traceln prints to the standard logger with the Trace level.
func (me *Logger) Traceln(v ...interface{}) {
	me.println(Ptrace, v...)
}
