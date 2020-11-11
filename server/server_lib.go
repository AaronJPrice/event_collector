package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func startServer(port int) *http.ServeMux {
	handler := http.NewServeMux()
	go listenAndServe(port, handler)
	return handler
}

func listenAndServe(port int, handler *http.ServeMux) {
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), handler))
}

func logRequest(r *http.Request) {
	log.Printf("Received: %+v", r)
}

func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Println("[warning] bad request: ", err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func methodNotAllowedResponse(w http.ResponseWriter, r *http.Request, allowedMethods string) {
	log.Println("[warning] method not allowed: ", r.Method)
	w.Header().Set("Allow", allowedMethods)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func internalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Println("[error]: ", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

// Is this really the best way of getting the source net.IP from a http.Request?
func getIP(r *http.Request) (net.IP, error) {
	ipString, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return nil, err
	}
	return net.ParseIP(ipString), err
}

func getPathItems(r *http.Request, indexes ...int) ([]string, error) {
	chosen := make([]string, len(indexes))
	items := strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")

	for i, n := range indexes {
		if n < 0 || n >= len(items) {
			return nil, errors.New("index out of range") // TODO: conv to const error type
		}
		chosen[i] = items[n]
	}

	return chosen, nil
}
