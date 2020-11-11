package filestore

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Start starts the filestore
func Start(filename string) (FS, error) {
	file, err := os.Create(filename)
	if err != nil {
		return FS{}, fmt.Errorf("could not start file-handler: %v", err)
	}
	writeChan := make(chan fileStoreMsg)
	doneChan := make(chan interface{})
	go fileHandlerRoutine(file, writeChan, doneChan)
	fs := FS{
		writeChan,
		doneChan,
	}

	return fs, nil
}

// FS represents the file-store
type FS struct {
	writeChan chan<- fileStoreMsg
	doneChan  chan interface{}
}

func (f FS) Write(data []byte) error {
	errChan := make(chan error)
	msg := fileStoreMsg{data: data, errChan: errChan}
	select {
	case f.writeChan <- msg:
		return <-errChan
	case <-f.doneChan:
		return errors.New("store closed") //TODO: Convert to const error type
	}
}

// Stop stops the file-store
func (f FS) Stop() {
	close(f.doneChan)
}

func fileHandlerRoutine(file *os.File, writeChan <-chan fileStoreMsg, doneChan chan interface{}) {
	defer file.Close()
	for {
		select {
		case msg := <-writeChan:
			_, err := file.Write(msg.data)
			msg.errChan <- err
		case <-doneChan:
			log.Println("fileHandlerRoutine shutting down.")
			return
		}
	}
}

type fileStoreMsg struct {
	data    []byte
	errChan chan<- error
}
