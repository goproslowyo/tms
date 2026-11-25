package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

var macRe = regexp.MustCompile(`^[0-9a-f]{12}$`)

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}

// getMacFromRequest returns the MAC from the "mac" query parameter as 12
// lowercase hex digits. Anything else is rejected, which also rules out
// path traversal.
func getMacFromRequest(r *http.Request) (string, error) {
	raw := r.URL.Query().Get("mac")
	log.Debugf("Raw MAC: %s", raw)
	mac := strings.ToLower(strings.NewReplacer(":", "", "-", "").Replace(raw))
	if !macRe.MatchString(mac) {
		return "", fmt.Errorf("invalid mac %q", raw)
	}
	log.Debugf("Sanitized MAC %s", mac)
	return mac, nil
}
