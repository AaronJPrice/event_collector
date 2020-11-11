package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"event_collector/objects"
)

// EventBody defines the body of a request sent the the event collector
type EventBody struct {
	PublisherID int
}

type store interface {
	Write([]byte) error
}

// StartMain starts the event collector server
func StartMain(port int, s store) {

	// Bind the store that has been passed in, which only knows how to handle binary data,
	// into a function which converts the event into binary data.
	// We create this function here and pass it in so that the HTTP handler and parse code can be tested
	// independently of the store code, by passing in a mocked store.
	eventStoreFunc := func(e objects.Event) error {
		bytes, err := e.FileFormat()
		if err != nil {
			return err
		}

		return s.Write(bytes)
	}

	handler := startServer(port)
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { handleEvent(w, r, eventStoreFunc) })
}

func handleEvent(w http.ResponseWriter, r *http.Request, eventStoreFunc func(objects.Event) error) {
	logRequest(r)
	switch r.Method {
	case http.MethodPut:
		eventData, err := parseRequest(r)
		if err != nil {
			badRequestResponse(w, r, err)
			return
		}

		err = eventStoreFunc(eventData)
		if err != nil {
			internalServerErrorResponse(w, r, err)
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		methodNotAllowedResponse(w, r, http.MethodPut)

	}
}

// Note: No validation here that the publisherid is in the body
// unmarshalling an empty json will set publisher id to 0
func parseRequest(r *http.Request) (objects.Event, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return objects.Event{}, fmt.Errorf("read body: %v", err)
	}

	ed := EventBody{}
	err = json.Unmarshal(body, &ed)
	if err != nil {
		return objects.Event{}, fmt.Errorf("unmarshall JSON: %v", err)
	}

	ip, err := getIP(r)
	if err != nil {
		return objects.Event{}, fmt.Errorf("get source IP: %v", err)
	}

	d := objects.Event{
		TS:          time.Now(),
		UserAgent:   r.UserAgent(),
		IP:          ip,
		PublisherID: ed.PublisherID,
		PageURL:     r.URL.String(),
	}

	return d, nil
}

func storeEvent(s store, e objects.Event) error {
	bytes, err := e.FileFormat()
	if err != nil {
		return err
	}

	return s.Write(bytes)
}
