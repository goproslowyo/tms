package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func metadataHandler(w http.ResponseWriter, r *http.Request) {
	// Parse MAC from request
	mac, err := getMacFromRequest(r)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Errorf("Error getting MAC from request: %s. Request: %+v", err, r)
		return
	}

	// Load configuration from YAML file
	configPath := fmt.Sprintf("%s/%s.yaml", CONFIG_DIR, mac)
	log.Printf("Attempting to load from %s", configPath)
	config, err := loadConfig(configPath)

	// If the file doesn't exist, return a 404
	if os.IsNotExist(err) {
		http.Error(w, "Not Found", http.StatusNotFound)
		log.Errorf("Config file not found: %s", err)
		return
	}

	// Handle other errors
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Errorf("Error loading config: %s", err)
		return
	}

	// Return the config as YAML
	w.Header().Set("Content-Type", "application/x-yaml")
	// Write the YAML to the response
	out, err := yaml.Marshal(config)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Errorf("Error marshalling config: %s", err)
		return
	}
	w.Write(out)
	// Close the connection
	w.(http.Flusher).Flush()
}
