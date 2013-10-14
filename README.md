loglevel [![Build Status](https://travis-ci.org/llimllib/loglevel.png)](https://travis-ci.org/llimllib/loglevel)
========

This project auns to be the simplest and best-tested levelled logging
wrapper around go's standard library log module. It retains as much of
the same API as possible.

Installation
============

`go get -u github.com/llimllib/loglevel`

Documentation
=============

You can find the documentation over at [GoDoc](http://godoc.org/github.com/llimllib/loglevel)

History
=======

Originally derived from [oneslang's log library](https://github.com/oneslang/log) but significantly
modified by Bill Mill.

Example
======
```go
package main

import (
	"bytes"
	"fmt"
	log "github.com/llimllib/loglevel"
)

func main() {
	// Set output level to info
	log.SetPriorityString("info")

    log.Info("Which means this will get printed")
	log.Warn("As will this")
    log.Debug("But not this")

	log.Info("The possible levels, in order from low to high:")
	log.Info("trace, debug, info, warn, error, fatal")

	// Just like the log module, you can set a prefix
	log.SetPrefix("OMG A PREFIX ")
	log.Info("Man that prefix is probably annoying")
	log.SetPrefix("")

	// You can also change what info is output in the prefix of each log msg
	log.SetFlags(log.Lshortfile | log.Lpriority)
	log.Info("Possible flags: Ldate, Ltime, Lmicroseconds, Llongfile")
	log.Info("                Lshortfile, Lpriority, LstdFlags = Ldate | Ltime")

	// Each log level has a format version and an ln version
	str := "this"
	log.Infof("Like %s", str)
	log.Errorln("Or this")

	// You can also make a log object; create a logger that outputs to buf, has
	// no prefix, outputs the standard line info, and prints logs at level info
	// and above
	buf := new(bytes.Buffer)
	l := log.New(buf, "", log.LstdFlags, log.Pinfo)
	l.Warn("This is a warning")
	fmt.Print(buf.String())

	// And, finally, you can cause your program to exit with Fatal, Fatalf, or
	// Fatalln
	log.Fatal("the program will terminate here")
	log.Info("so this line will not get printed")
}
```
