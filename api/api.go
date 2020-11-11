package api

import (
	"fmt"
	"log"

	"event_collector/filestore"
	"event_collector/server"
)

// Start starts the event collector application
func Start(port int, managementPort int, filename string) {
	log.Printf("Starting: event_collector. port: %v, managementPort %v", port, managementPort)

	fs, err := filestore.Start(filename)
	if err != nil {
		log.Fatal(fmt.Errorf("could not start store: %v", err))
	}
	server.StartManagement(managementPort)
	server.StartMain(port, fs)

	log.Println("Complete: event_collector.")
}
