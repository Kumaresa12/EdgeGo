package config

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLoad(t *testing.T) {

	dir := t.TempDir()

	file := filepath.Join(dir, "edgego.yaml")

	content := `
server:
  host: "localhost"
  port: 8080
  readTimeout: 10s
  writeTimeout: 10s
  idleTimeout: 60s

logging:
  level: "debug"
`

	err := os.WriteFile(
		file,
		[]byte(content),
		0644,
	)

	if err != nil {
		t.Fatalf("failed to create test config: %v", err)
	}

	cfg, err := Load(file)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.Server.Port != 8080 {
		t.Fatalf(
			"expected port 8080, got %d",
			cfg.Server.Port,
		)
	}

	if cfg.Server.ReadTimeout != 10*time.Second {
		t.Fatalf(
			"expected read timeout 10s, got %v",
			cfg.Server.ReadTimeout,
		)
	}
}

func TestLoadInvalidConfig(t *testing.T) {

	dir := t.TempDir()

	file := filepath.Join(
		dir,
		"invalid.yaml",
	)

	content := `
server:
  host: [
`

	err := os.WriteFile(
		file,
		[]byte(content),
		0644,
	)

	if err != nil {
		t.Fatalf("failed to create test config: %v", err)
	}

	_, err = Load(file)

	if err == nil {
		t.Fatal("expected config loading to fail")
	}
}
