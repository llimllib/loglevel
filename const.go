// Copyright 2012 The Oneslang Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

// Priority used for identifying the severity of an event.
const (
	Poff = iota
	Pfatal
	Perror
	Pwarn
	Pinfo
	Pdebug
	Ptrace
	Pall
)

var priorityName = []string{
	Poff:       "OFF",
	Pfatal:     "FATAL",
	Perror:     "ERROR",
	Pwarn:      "WARN",
	Pinfo:      "INFO",
	Pdebug:     "DEBUG",
	Ptrace:     "TRACE",
	Pall:       "ALL",
}

// Layouts used for identifying the format of an event.
const (
	// Debug 2012/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate      = 1 << iota     // the date: 2012/01/23
	Ltime                      // the time: 01:23:23
	Lmicroseconds              // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                  // full file name and line number: /a/b/c/d.go:23
	Lshortfile                 // final file name element and line number: d.go:23. overrides Llongfile
	Lpriority                  // the priority: Debug
	Lstd       = Ldate | Ltime // initial values for the standard logger
)
