package filestore

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"event_collector/test"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

//==============================================================================
// Tests
//==============================================================================
func TestWrite(t *testing.T) {
	expected := []byte("Some test data")
	filename := "TestWriteToFile.txt"

	fs, err := Start(filename)
	test.FatalErrCheck(t, err)

	fs.Write(expected)

	file, err := os.Open(filename)
	test.FatalErrCheck(t, err)

	actual, err := ioutil.ReadAll(file)
	test.FatalErrCheck(t, err)

	if !bytes.Equal(actual, expected) {
		t.Fatal("Actual:", actual, "Expected:", expected)
	}

	fs.Stop()
	err = os.Remove(filename)
	test.FatalErrCheck(t, err)
}

//==============================================================================
// Benchmarks
//==============================================================================
func BenchmarkWrite(b *testing.B) {
	expected := []byte("Some test data")
	filename := "TestWriteToFile.txt"

	fs, err := Start(filename)
	test.FatalErrCheck(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fs.Write(expected)
	}
	b.StopTimer()

	fs.Stop()
	err = os.Remove(filename)
	test.FatalErrCheck(b, err)
}
