// Copyright 2012 The Oneslang Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Layouts used for identifying the format of an event.
package layout

const (
	// Debug 2012/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Date         = 1 << iota   // the date: 2012/01/23
	Time                       // the time: 01:23:23
	Microseconds               // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Longfile                   // full file name and line number: /a/b/c/d.go:23
	Shortfile                  // final file name element and line number: d.go:23. overrides Llongfile
	Level                      // the level: Debug
	Default      = Date | Time // initial values for the standard logger
)
