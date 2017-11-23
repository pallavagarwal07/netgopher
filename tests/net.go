package main

import (
	"testing"

	"github.com/pallavagarwal07/gophernet"
)

// Copied from testing.TB: The testing.TB version has a private
// method in the signature. This TB will always thus be a subset
// of testing.TB interface (tested statically below).
type TB interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
	Helper()
}

// Static assertion
var _ TB = testing.TB(nil)

// We can't use the standard Go test toolchain because we need to run the same
// tests with phantomjs + gopherjs compiled js files. Renaming this file with
// *_test.go format causes the `go run` to refuse to compile the test_runner.go
// (used for the js tests).
var tests = map[string]func(t TB){
	"TestGet":  TestGet,
	"TestPost": TestPost,
}

func TestGet(t TB) {
	_, err := gophernet.Get("https://raw.githubusercontent.com/kovidgoyal/kitty/master/.travis.yml", nil)
	if err != nil {
		t.Fatalf("Get failed with error %v", err)
	}
}

func TestPost(t TB) {
	t.Fail()
}
