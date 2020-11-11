package test

import (
	"path/filepath"
	"runtime"
	"testing"
)

// Assert checks if the values are equal and calls t.Fatal with a useful message if not
func Assert(t *testing.T, value interface{}, expected interface{}, context ...string) {
	_, file, ln, _ := runtime.Caller(1)
	if value != expected {
		t.Logf("Assert failed on %v line %v", filepath.Base(file), ln)
		if context != nil {
			t.Log(context)
		}
		t.Fatal("Value:", value, "Expected:", expected)
	}
}

// type equaler interface {
// 	Equals(interface{}) bool
// }

// // AssertEqual checks if the values are equal (for types with a .Equals method) and calls t.Fatal with a useful message if not
// func AssertEqual(t *testing.T, value equaler, expected equaler, context ...string) {
// 	_, file, ln, _ := runtime.Caller(1)
// 	if !value.Equals(expected) {
// 		t.Logf("Assert failed on %v line %v", filepath.Base(file), ln)
// 		if context != nil {
// 			t.Log(context)
// 		}
// 		t.Fatal("Value:", value, "Expected:", expected)
// 	}
// }

type testObject interface {
	Fatal(...interface{})
}

// FatalErrCheck calls t.Fatal if err is not nil
func FatalErrCheck(t testObject, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
