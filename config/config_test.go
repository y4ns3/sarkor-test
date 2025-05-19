package config

import (
	"github.com/joho/godotenv"
	"testing"
)

func TestNewConfig(t *testing.T) {

	err := godotenv.Load(".env.test")
	if err != nil {
		t.Error("Error loading .env file")
	}
	cfg, err := NewConfig()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	expectedUrl := "postgres://test:test@localhost:5432/test"
	if cfg.Db.Url != expectedUrl {
		t.Errorf("expected Db.Url %s, got %s", expectedUrl, cfg.Db.Url)
	}

	if cfg.Server.Port != "8080" || cfg.Server.Host != "127.0.0.1" {
		t.Errorf("expected server host 127.0.0.1 and port 8080, got host %s port %s", cfg.Server.Host, cfg.Server.Port)
	}
}
