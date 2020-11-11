package main

import (
	"flag"

	"event_collector/api"
)

func main() {
	port := flag.Int("p", 80, "Port that server will listen on.")
	managementPort := flag.Int("m", 700, "Port that server will listen on for application monitoring and management.")
	eventLogFilename := flag.String("f", "eventLog", "Name of file to write events to.")
	flag.Parse()

	api.Start(*port, *managementPort, *eventLogFilename)

	select {} //Block indefinitely to keep app alive
}
