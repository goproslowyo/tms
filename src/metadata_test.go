package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// A trimmed multi-document Talos config, with comments, in the shape
// `talosctl gen config` produces. It must come back byte-for-byte.
const talosConfig = `version: v1alpha1 # Indicates the schema used to decode the contents.
debug: false
machine:
    type: controlplane
    token: abc123.def456
    features:
        kubePrism:
            enabled: true
            port: 7445
    registries:
        mirrors:
            docker.io:
                endpoints:
                    - http://10.0.0.2:5000
cluster:
    clusterName: test
---
apiVersion: v1alpha1
kind: HostnameConfig
hostname: cp1
auto: "off"
---
apiVersion: v1alpha1
kind: UnattendedInstallConfig
provisioning:
    diskSelector:
        match: disk.dev_path == "/dev/sda"
`

func setupConfigDir(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "aabbcc001122.yaml"), []byte(talosConfig), 0o644); err != nil {
		t.Fatal(err)
	}
	// A file outside the config directory that traversal must never reach.
	if err := os.WriteFile(filepath.Join(filepath.Dir(dir), "secret.yaml"), []byte("nope"), 0o644); err != nil {
		t.Fatal(err)
	}
	old := CONFIG_DIR
	CONFIG_DIR = dir
	t.Cleanup(func() { CONFIG_DIR = old })
}

func serve(method, target string) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	metadataHandler(rec, httptest.NewRequest(method, target, nil))
	return rec
}

func TestServesConfigByteForByte(t *testing.T) {
	setupConfigDir(t)
	// Every MAC form, including the URL-encoded colons Talos sends.
	for _, q := range []string{
		"mac=aa:bb:cc:00:11:22",
		"mac=AA:BB:CC:00:11:22",
		"mac=aa%3Abb%3Acc%3A00%3A11%3A22",
		"mac=aa-bb-cc-00-11-22",
		"mac=aabbcc001122",
	} {
		rec := serve(http.MethodGet, "/talos/config?"+q)
		if rec.Code != http.StatusOK {
			t.Fatalf("%s: status %d, want 200", q, rec.Code)
		}
		if got := rec.Body.String(); got != talosConfig {
			t.Fatalf("%s: body differs from the stored file:\n%s", q, got)
		}
		if ct := rec.Header().Get("Content-Type"); ct != "application/x-yaml" {
			t.Fatalf("%s: Content-Type %q", q, ct)
		}
	}
}

func TestHead(t *testing.T) {
	setupConfigDir(t)
	if rec := serve(http.MethodHead, "/talos/config?mac=aabbcc001122"); rec.Code != http.StatusOK {
		t.Fatalf("HEAD status %d, want 200", rec.Code)
	}
}

func TestErrors(t *testing.T) {
	setupConfigDir(t)
	cases := []struct {
		name, method, target string
		want                 int
	}{
		{"unknown mac", http.MethodGet, "/talos/config?mac=aa:bb:cc:00:11:99", http.StatusNotFound},
		{"missing mac", http.MethodGet, "/talos/config", http.StatusBadRequest},
		{"empty mac", http.MethodGet, "/talos/config?mac=", http.StatusBadRequest},
		{"too short", http.MethodGet, "/talos/config?mac=aabbcc", http.StatusBadRequest},
		{"not hex", http.MethodGet, "/talos/config?mac=zzbbcc001122", http.StatusBadRequest},
		{"traversal", http.MethodGet, "/talos/config?mac=../secret", http.StatusBadRequest},
		{"encoded traversal", http.MethodGet, "/talos/config?mac=..%2Fsecret", http.StatusBadRequest},
		{"other path", http.MethodGet, "/other?mac=aabbcc001122", http.StatusNotFound},
		{"root path", http.MethodGet, "/?mac=aabbcc001122", http.StatusNotFound},
		{"post", http.MethodPost, "/talos/config?mac=aabbcc001122", http.StatusMethodNotAllowed},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rec := serve(c.method, c.target)
			if rec.Code != c.want {
				t.Fatalf("status %d, want %d", rec.Code, c.want)
			}
			if rec.Body.String() == talosConfig || rec.Body.String() == "nope" {
				t.Fatal("served file contents on an error response")
			}
		})
	}
}

func TestGetMacFromRequest(t *testing.T) {
	for raw, want := range map[string]string{
		"aa:bb:cc:dd:ee:ff": "aabbccddeeff",
		"AA-BB-CC-DD-EE-FF": "aabbccddeeff",
		"AaBbCcDdEeFf":      "aabbccddeeff",
	} {
		r := httptest.NewRequest(http.MethodGet, "/talos/config", nil)
		q := r.URL.Query()
		q.Set("mac", raw)
		r.URL.RawQuery = q.Encode()
		got, err := getMacFromRequest(r)
		if err != nil || got != want {
			t.Fatalf("%q -> %q, %v; want %q", raw, got, err, want)
		}
	}
}
