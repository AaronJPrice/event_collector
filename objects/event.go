package objects

import (
	"encoding/json"
	"net"
	"time"
)

// Event is a struct which represents an event
type Event struct {
	TS          time.Time
	UserAgent   string
	IP          net.IP
	PublisherID int
	PageURL     string
}

// JSON converts an event to JSON format
func (e Event) JSON() ([]byte, error) {
	return json.Marshal(e)
}

// FileFormat converts an event to format suitable for file storage
func (e Event) FileFormat() ([]byte, error) {
	bytes, err := e.JSON()
	if err != nil {
		return nil, err
	}

	fileFormat := append(bytes, []byte("\n")...) // Add a newline at end so it writes nicely

	return fileFormat, nil
}
