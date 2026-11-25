package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
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
	server := &http.Server{
		Addr:              listen,
		Handler:           http.HandlerFunc(metadataHandler),
		ReadHeaderTimeout: 10 * time.Second,
	}
	log.Printf("Listening on %s", listen)
	log.Fatal(server.ListenAndServe())
}
