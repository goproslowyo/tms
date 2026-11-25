package main

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// metadataHandler serves the machine config for the MAC in the request
// byte-for-byte. Talos configs are multi-document YAML whose schema changes
// between releases, so the file is never parsed or re-encoded here.
func metadataHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/talos/config" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	mac, err := getMacFromRequest(r)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Errorf("Invalid MAC in request from %s: %s", r.RemoteAddr, err)
		return
	}

	configPath := filepath.Join(CONFIG_DIR, mac+".yaml")
	log.Printf("Attempting to load from %s", configPath)
	data, err := os.ReadFile(configPath)
	if errors.Is(err, fs.ErrNotExist) {
		http.Error(w, "Not Found", http.StatusNotFound)
		log.Errorf("Config file not found: %s", err)
		return
	}
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Errorf("Error loading config: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/x-yaml")
	if _, err := w.Write(data); err != nil {
		log.Errorf("Error writing response: %s", err)
	}
}
