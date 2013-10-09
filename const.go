package loglevel

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
	Poff:   "OFF",
	Pfatal: "FATAL",
	Perror: "ERROR",
	Pwarn:  "WARN",
	Pinfo:  "INFO",
	Pdebug: "DEBUG",
	Ptrace: "TRACE",
	Pall:   "ALL",
}

// Flags used for identifying the format of an event.
const (
	// Bits or'ed together to control what's printed. There is no control over the
	// order they appear (the order listed here) or the format they present (as
	// described in the comments).  A colon appears after these items:
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // the date: 2012/01/23
	Ltime                         // the time: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	Lpriority                     // the priority: Debug
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
