package gocql

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
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
	capture bytes.Buffer
}

func (l *testLogger) Print(v ...interface{})                 { fmt.Fprint(&l.capture, v...) }
func (l *testLogger) Printf(format string, v ...interface{}) { fmt.Fprintf(&l.capture, format, v...) }
func (l *testLogger) Println(v ...interface{})               { fmt.Fprintln(&l.capture, v...) }
func (l *testLogger) String() string                         { return l.capture.String() }

type defaultLogger struct{}

func (l *defaultLogger) Print(v ...interface{}) {
	args := append([]any{getFileAndLine()}, v...)
	log.Print(args...)
}
func (l *defaultLogger) Printf(format string, v ...interface{}) {
	log.Printf(getFileAndLine()+" "+format, v...)
}
func (l *defaultLogger) Println(v ...interface{}) {
	args := append([]any{getFileAndLine()}, v...)
	log.Println(args...)
}

func getFileAndLine() string {
	_, filename, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d:", filename, line)
}

// Logger for logging messages.
// Deprecated: Use ClusterConfig.Logger instead.
var Logger StdLogger = &defaultLogger{}
