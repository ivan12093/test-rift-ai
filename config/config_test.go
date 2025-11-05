package config

import (
	"os"
	"testing"
)

func TestLoadServerConfig(t *testing.T) {
	// Save original values
	originalPort := os.Getenv("PORT")
	originalDifficulty := os.Getenv("DIFFICULTY")

	// Clear environment variables
	os.Unsetenv("PORT")
	os.Unsetenv("DIFFICULTY")

	// Test default values
	cfg := LoadServerConfig()
	if cfg.ServerPort != "8080" {
		t.Errorf("LoadServerConfig() ServerPort = %v, want 8080", cfg.ServerPort)
	}
	if cfg.Difficulty != 20 {
		t.Errorf("LoadServerConfig() Difficulty = %v, want 20", cfg.Difficulty)
	}
	if cfg.QuotesFile != "quotes.txt" {
		t.Errorf("LoadServerConfig() QuotesFile = %v, want quotes.txt", cfg.QuotesFile)
	}

	// Test with environment variables
	os.Setenv("PORT", "9090")
	os.Setenv("DIFFICULTY", "16")
	cfg = LoadServerConfig()
	if cfg.ServerPort != "9090" {
		t.Errorf("LoadServerConfig() ServerPort = %v, want 9090", cfg.ServerPort)
	}
	if cfg.Difficulty != 16 {
		t.Errorf("LoadServerConfig() Difficulty = %v, want 16", cfg.Difficulty)
	}

	// Restore original values
	if originalPort != "" {
		os.Setenv("PORT", originalPort)
	} else {
		os.Unsetenv("PORT")
	}
	if originalDifficulty != "" {
		os.Setenv("DIFFICULTY", originalDifficulty)
	} else {
		os.Unsetenv("DIFFICULTY")
	}
}

func TestLoadClientConfig(t *testing.T) {
	// Save original value
	originalAddr := os.Getenv("SERVER_ADDR")

	// Clear environment variable
	os.Unsetenv("SERVER_ADDR")

	// Test default value
	cfg := LoadClientConfig()
	if cfg.ServerAddr != "server:8080" {
		t.Errorf("LoadClientConfig() ServerAddr = %v, want server:8080", cfg.ServerAddr)
	}

	// Test with environment variable
	os.Setenv("SERVER_ADDR", "localhost:9090")
	cfg = LoadClientConfig()
	if cfg.ServerAddr != "localhost:9090" {
		t.Errorf("LoadClientConfig() ServerAddr = %v, want localhost:9090", cfg.ServerAddr)
	}

	// Restore original value
	if originalAddr != "" {
		os.Setenv("SERVER_ADDR", originalAddr)
	} else {
		os.Unsetenv("SERVER_ADDR")
	}
}
