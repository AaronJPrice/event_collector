package objects

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

//==============================================================================
// Benchmarks
//==============================================================================
func BenchmarkJSON(b *testing.B) {
	e := newEvent()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.JSON()
	}
}

func BenchmarkFileFormat(b *testing.B) {
	e := newEvent()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.FileFormat()
	}
}

//==============================================================================
// Utilities
//==============================================================================
func newEvent() Event {
	return Event{
		TS:          time.Now(),
		UserAgent:   "UserAgent",
		IP:          net.ParseIP("127.0.0.1"),
		PublisherID: 12345,
		PageURL:     "PageURL",
	}
}
