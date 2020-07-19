package config

import "testing"

func TestLoad(t *testing.T) {
	Load()
	cfg := Get()

	if cfg.Server.Host != "localhost:8080" ||
		cfg.Server.WriteTimeout != 5 {
		t.Fatalf("Configuration not as expected: %v", cfg)
	}
}
