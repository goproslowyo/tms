package main

import (
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}

func getMacFromRequest(r *http.Request) (string, error) {
	// Get url parameters
	mac := strings.ToLower(r.URL.Query().Get("mac"))
	log.Debugf("Raw MAC: %s", mac)
	// Sanitize input to prevent directory traversal \/.. etc.
	mac = strings.Replace(mac, "/", "", -1)
	mac = strings.Replace(mac, "\\", "", -1)
	mac = strings.Replace(mac, "..", "", -1)
	mac = strings.Replace(mac, ":", "", -1)
	log.Debugf("Sanitized MAC %s", mac)
	return mac, nil
}

func loadConfig(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
