package loglevel

import (
	"bytes"
	"os"
	"regexp"
	"testing"
)

const (
	Rdate         = `[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9]`
	Rtime         = `[0-9][0-9]:[0-9][0-9]:[0-9][0-9]`
	Rmicroseconds = `\.[0-9][0-9][0-9][0-9][0-9][0-9]`
	Rline         = `(55|63):` // must update if the calls to l.Printf / l.Print below move
	Rlongfile     = `.*/[A-Za-z0-9_\-]+\.go:` + Rline
	Rshortfile    = `[A-Za-z0-9_\-]+\.go:` + Rline
	Rinfo         = `INFO `
)

type tester struct {
	flag    int
	prefix  string
	pattern string // regexp that log output must match; we add ^ and expected_text$ always
}

var tests = []tester{
	// individual pieces:
	{0, "", ""},
	{0, "XXX", "XXX"},
	{Ldate, "", Rdate + " "},
	{Ltime, "", Rtime + " "},
	{Ltime | Lmicroseconds, "", Rtime + Rmicroseconds + " "},
	{Lmicroseconds, "", Rtime + Rmicroseconds + " "}, // microsec implies time
	{Llongfile, "", Rlongfile + " "},
	{Lshortfile, "", Rshortfile + " "},
	{Llongfile | Lshortfile, "", Rshortfile + " "}, // shortfile overrides longfile
	{Lpriority, "", Rinfo},
	{Lpriority | Ltime, "", Rinfo + Rtime + " "},
	// everything at once:
	{Ldate | Ltime | Lmicroseconds | Llongfile, "XXX", "XXX" + Rdate + " " + Rtime + Rmicroseconds + " " + Rlongfile + " "},
	{Ldate | Ltime | Lmicroseconds | Lshortfile, "XXX", "XXX" + Rdate + " " + Rtime + Rmicroseconds + " " + Rshortfile + " "},
}

// Test using Println("hello", 23, "world") or using Printf("hello %d world", 23)
func testPrint(t *testing.T, flag int, prefix string, pattern string, useFormat bool) {
	buf := new(bytes.Buffer)
	SetOutput(buf)
	SetFlags(flag)
	SetPrefix(prefix)
	SetPriority(Pall)
	if useFormat {
		Infof("hello %d world", 23)
	} else {
		Infoln("hello", 23, "world")
	}
	line := buf.String()
	line = line[0 : len(line)-1]
	pattern = "^" + pattern + "hello 23 world$"
	matched, err := regexp.MatchString(pattern, line)
	if err != nil {
		t.Fatal("pattern did not compile:", err)
	}
	if !matched {
		t.Errorf("log output should match %q but is %q", pattern, line)
	}
	SetOutput(os.Stderr)
}

func TestAll(t *testing.T) {
	for _, testcase := range tests {
		testPrint(t, testcase.flag, testcase.prefix, testcase.pattern, false)
		testPrint(t, testcase.flag, testcase.prefix, testcase.pattern, true)
	}
}

func TestOutput(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "", 0, Pdebug)
	l.Infoln(testString)
	if expect := testString + "\n"; b.String() != expect {
		t.Errorf("log output should match %q is %q", expect, b.String())
	}
}

func TestFlagAndPrefixSetting(t *testing.T) {
	var b bytes.Buffer
	l := New(&b, "Test:", LstdFlags, Pdebug)
	f := l.Flags()
	if f != LstdFlags {
		t.Errorf("Flags 1: expected %x got %x", LstdFlags, f)
	}
	l.SetFlags(f | Lmicroseconds)
	f = l.Flags()
	if f != LstdFlags|Lmicroseconds {
		t.Errorf("Flags 2: expected %x got %x", LstdFlags|Lmicroseconds, f)
	}
	p := l.Prefix()
	if p != "Test:" {
		t.Errorf(`Prefix: expected "Test:" got %q`, p)
	}
	l.SetPrefix("Reality:")
	p = l.Prefix()
	if p != "Reality:" {
		t.Errorf(`Prefix: expected "Reality:" got %q`, p)
	}
	// Verify a log message looks right, with our prefix and microseconds present.
	l.Info("hello")
	pattern := "^Reality:" + Rdate + " " + Rtime + Rmicroseconds + " hello\n"
	matched, err := regexp.Match(pattern, b.Bytes())
	if err != nil {
		t.Fatalf("pattern %q did not compile: %s", pattern, err)
	}
	if !matched {
		t.Error("message did not match pattern")
	}
}

func TestPriorityString(t *testing.T) {
	buf := new(bytes.Buffer)
	l := New(buf, "", 0, Pinfo)
	if l.Priority() != Pinfo {
		t.Fatalf("Priority should be Pinfo but is %s", priorityName[l.Priority()])
	}

	err := l.SetPriorityString("WARN")
	if err != nil {
		t.Fatalf("SetPriorityString returned an error: %s", err)
	}
	if l.Priority() != Pwarn {
		t.Fatalf("Priority should be Pwarn but is %s", priorityName[l.Priority()])
	}

	err = l.SetPriorityString("debug")
	if err != nil {
		t.Fatalf("SetPriorityString returned an error: %s", err)
	}
	if l.Priority() != Pdebug {
		t.Fatalf("Priority should be Pdebug but is %s", priorityName[l.Priority()])
	}

	err = l.SetPriorityString("does not exist")
	if err == nil {
		t.Fatal("SetPriorityString ought to have errored on invalid input")
	}
	if l.Priority() != Pdebug {
		t.Fatalf("Priority should be Pdebug but is %s", priorityName[l.Priority()])
	}
}

func testPriorityLevel(t *testing.T, testlevel int, f func(v ...interface{})) {
	for level := range priorityName {
		buf := new(bytes.Buffer)
		SetOutput(buf)
		SetPriority(level)
		SetFlags(0)
		SetPrefix("")
		f("a")
		if level >= testlevel {
			if buf.String() != "a\n" {
				t.Fatalf("Expected 'a\\n' with level %s and testlevel %s, got <%s>", level, testlevel, buf.String())
			}
		} else {
			if len(buf.String()) != 0 {
				t.Fatalf("Expected '' with level %s and testlevel %s, got <%s>", level, testlevel, buf.String())
			}
		}
	}
}

var funcs = map[int]func(v ...interface{}){
	Ptrace: Trace,
	Pdebug: Debug,
	Pinfo:  Info,
	Pwarn:  Warn,
	Perror: Error,
}

func TestPriority(t *testing.T) {
	for k, v := range funcs {
		testPriorityLevel(t, k, v)
	}
}

func TestPriorityAll(t *testing.T) {
	for _, f := range funcs {
		buf := new(bytes.Buffer)
		SetOutput(buf)
		SetPriority(Pall)
		SetFlags(0)
		SetPrefix("")
		f("a")
		if buf.String() != "a\n" {
			t.Fatalf("Expected 'a\\n' with func %s and testlevel Pall, got <%s>", f, buf.String())
		}
	}
}

func TestPriorityOff(t *testing.T) {
	for _, f := range funcs {
		buf := new(bytes.Buffer)
		SetOutput(buf)
		SetPriority(Poff)
		SetFlags(0)
		SetPrefix("")
		f("a")
		if buf.String() != "" {
			t.Fatalf("Expected '' with func %s and testlevel Poff, got <%s>", f, buf.String())
		}
	}
}

// XXX How to test fatal and panic?
//func TestFatal(t *testing.T) {
//}
