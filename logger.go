package gocql

import (
	"bytes"
	"fmt"
	"log"
	"sync"
)

type StdLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type nopLogger struct{}

func (n nopLogger) Print(_ ...interface{}) {}

func (n nopLogger) Printf(_ string, _ ...interface{}) {}

func (n nopLogger) Println(_ ...interface{}) {}

type testLogger struct {
	mu      sync.Mutex
	capture bytes.Buffer
}

// Write appends the contents of p to the buffer, growing the buffer as needed. It returns
// the number of bytes written.
func (l *testLogger) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.capture.Write(p)
}

func (l *testLogger) String() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.capture.String()
}

func (l *testLogger) Print(v ...interface{})                 { fmt.Fprint(l, v...) }
func (l *testLogger) Printf(format string, v ...interface{}) { fmt.Fprintf(l, format, v...) }
func (l *testLogger) Println(v ...interface{})               { fmt.Fprintln(l, v...) }

//func (l *testLogger) String() string                         { return l.String() }

type defaultLogger struct{}

func (l *defaultLogger) Print(v ...interface{})                 { log.Print(v...) }
func (l *defaultLogger) Printf(format string, v ...interface{}) { log.Printf(format, v...) }
func (l *defaultLogger) Println(v ...interface{})               { log.Println(v...) }

// Logger for logging messages.
// Deprecated: Use ClusterConfig.Logger instead.
var Logger StdLogger = &defaultLogger{}
