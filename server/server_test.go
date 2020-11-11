package server

import (
	"net"
	"net/http"
	"testing"

	"event_collector/test"
)

//==============================================================================
// Tests
//==============================================================================
func TestGetIP(t *testing.T) {
	ip := "127.0.0.1"
	expected := net.ParseIP(ip)
	r := &http.Request{RemoteAddr: ip + ":111"}

	actual, err := getIP(r)
	test.FatalErrCheck(t, err)

	test.Assert(t, actual.String(), expected.String())
}

// func TestHandleEvent(t *testing.T) {
// 	testCh := make(chan objects.Event)

// 	mockEventStoreFunc := func(e objects.Event) error {
// 		testCh <- e
// 		return nil
// 	}

// 	mockWriter := false
// 	mockRequest := false

// 	handleEvent(mockWriter, mockRequest, mockEventStoreFunc)
// }

// type mockWriter struct{}
