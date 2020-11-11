package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"event_collector/server"
)

var client = &http.Client{}

// Status gets the app status.
func Status(host string, port int) (resp *http.Response, err error) {
	return request(http.MethodGet, url(host, port, "/status"), nil)
}

// Enable sets app status to enabled.
func Enable(host string, port int) (resp *http.Response, err error) {
	return request(http.MethodPost, url(host, port, "/status/enable"), nil)
}

// Disable sets app status to disabled.
func Disable(host string, port int) (resp *http.Response, err error) {
	return request(http.MethodPost, url(host, port, "/status/disable"), nil)
}

// Event sends an event to the app.
func Event(host string, port int, path string, publisherID int) (resp *http.Response, err error) {
	ed := server.EventBody{PublisherID: publisherID}
	body, err := json.Marshal(ed)
	if err != nil {
		return nil, err
	}

	return request(http.MethodPut, url(host, port, path), bytes.NewBuffer(body))
}

func request(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func url(host string, port int, path string) string {
	return "http://" + host + ":" + strconv.Itoa(port) + path
}
