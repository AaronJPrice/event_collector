package server

import (
	"fmt"
	"net/http"
)

const (
	enable  = "enable"
	disable = "disable"
)

var enabled = false

// StartManagement starts the management server
func StartManagement(debugPort int) {
	debugHandler := startServer(debugPort)
	debugHandler.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) { statusHandler(w, r) })
	debugHandler.HandleFunc("/status/"+enable, func(w http.ResponseWriter, r *http.Request) { newHandler(w, r) })
	debugHandler.HandleFunc("/status/"+disable, func(w http.ResponseWriter, r *http.Request) { newHandler(w, r) })
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println()

	logRequest(r)
	switch r.Method {

	case http.MethodPost:
		pathItems, err := getPathItems(r, 1)
		if err != nil {
			internalServerErrorResponse(w, r, err)
		}
		switch pathItems[0] {
		case enable:
			enabled = true
		case disable:
			enabled = false
		}
		w.WriteHeader(http.StatusOK)

	default:
		methodNotAllowedResponse(w, r, http.MethodGet)

	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	switch r.Method {

	case http.MethodGet:
		switch enabled {
		case true:
			w.WriteHeader(http.StatusOK)
		case false:
			w.WriteHeader(http.StatusServiceUnavailable)
		}

	default:
		methodNotAllowedResponse(w, r, http.MethodGet)

	}
}

func enableHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	switch r.Method {

	case http.MethodPost:
		enabled = true
		w.WriteHeader(http.StatusOK)

	default:
		methodNotAllowedResponse(w, r, http.MethodGet)

	}
}

func disableHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	switch r.Method {

	case http.MethodPost:
		enabled = false
		w.WriteHeader(http.StatusOK)

	default:
		methodNotAllowedResponse(w, r, http.MethodGet)

	}
}
