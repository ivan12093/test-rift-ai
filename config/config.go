package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort   string
	Difficulty   int
	QuotesFile   string
	Timeout      int
	ServerAddr   string
}

func LoadServerConfig() *Config {
	port := getEnv("PORT", "8080")
	difficulty, _ := strconv.Atoi(getEnv("DIFFICULTY", "20"))
	quotesFile := getEnv("QUOTES_FILE", "quotes.txt")
	timeout, _ := strconv.Atoi(getEnv("TIMEOUT_SECONDS", "30"))

	return &Config{
		ServerPort: port,
		Difficulty: difficulty,
		QuotesFile: quotesFile,
		Timeout:    timeout,
	}
}

func LoadClientConfig() *Config {
	serverAddr := getEnv("SERVER_ADDR", "server:8080")

	return &Config{
		ServerAddr: serverAddr,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
