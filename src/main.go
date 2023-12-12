package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
)

// Global variables
var CONFIG_DIR string = "./configs"
var LISTEN string = "127.0.0.1:8088"
var LOG_LEVEL string = "info"

func main() {
	// Get log level from environment MDS_LOG_LEVEL variable if set.
	LOG_LEVEL = getEnv("MDS_LOG_LEVEL", LOG_LEVEL)
	logLevel, err := log.ParseLevel(LOG_LEVEL)
	if err != nil {
		log.Fatalf("Error parsing log level: %s", err)
	}
	log.SetLevel(logLevel)
	log.Debugf("Log level set to %s", logLevel)

	// Get listen interface and port from environment MDS_LISTEN variable
	listen := getEnv("MDS_LISTEN", LISTEN)
	log.Debugf("Listen interface: %s", listen)

	// Get config directory from environment MDS_CONFIG_DIR variable
	CONFIG_DIR = getEnv("MDS_CONFIG_DIR", CONFIG_DIR)
	log.Debugf("Config directory: %s", CONFIG_DIR)

	log.Println("Starting Talos Metadata Config Service")
	http.HandleFunc("/talos/config", metadataHandler)

	// Enable HTTP/2 support
	http2.VerboseLogs = true // Enable verbose logs for debugging
	// Create an HTTP/2 enabled server
	server := &http.Server{
		Addr:    listen,
		Handler: http.HandlerFunc(metadataHandler),
	}
	// Enable HTTP/2 support on the server
	http2.ConfigureServer(server, nil)
	log.Printf("Listening on %s", listen)
	server.ListenAndServe()
}
