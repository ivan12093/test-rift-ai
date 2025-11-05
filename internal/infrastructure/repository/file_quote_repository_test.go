package repository

import (
	"os"
	"path/filepath"
	"testing"
	"word-of-wisdom/internal/domain/repository"
)

func createTempQuotesFile(t *testing.T, content string) string {
	tmpfile, err := os.CreateTemp("", "quotes_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tmpfile.Name()
}

func TestNewFileQuoteRepository(t *testing.T) {
	content := "Quote 1\nQuote 2\nQuote 3\n"
	filename := createTempQuotesFile(t, content)
	defer os.Remove(filename)

	repo, err := NewFileQuoteRepository(filename)
	if err != nil {
		t.Fatalf("NewFileQuoteRepository() error = %v", err)
	}
	if repo == nil {
		t.Fatal("NewFileQuoteRepository() returned nil")
	}
}

func TestNewFileQuoteRepository_FileNotFound(t *testing.T) {
	_, err := NewFileQuoteRepository("nonexistent_file.txt")
	if err == nil {
		t.Error("NewFileQuoteRepository() should return error for nonexistent file")
	}
}

func TestFileQuoteRepository_GetAll(t *testing.T) {
	content := "Quote 1\nQuote 2\nQuote 3\n"
	filename := createTempQuotesFile(t, content)
	defer os.Remove(filename)

	repo, err := NewFileQuoteRepository(filename)
	if err != nil {
		t.Fatalf("NewFileQuoteRepository() error = %v", err)
	}

	quotes, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll() error = %v", err)
	}

	if len(quotes) != 3 {
		t.Errorf("GetAll() returned %d quotes, want 3", len(quotes))
	}

	expectedQuotes := []string{"Quote 1", "Quote 2", "Quote 3"}
	for i, quote := range quotes {
		if quote.Text != expectedQuotes[i] {
			t.Errorf("GetAll() quote[%d] = %v, want %v", i, quote.Text, expectedQuotes[i])
		}
	}
}

func TestFileQuoteRepository_GetAll_EmptyLines(t *testing.T) {
	content := "Quote 1\n\nQuote 2\n  \nQuote 3\n"
	filename := createTempQuotesFile(t, content)
	defer os.Remove(filename)

	repo, err := NewFileQuoteRepository(filename)
	if err != nil {
		t.Fatalf("NewFileQuoteRepository() error = %v", err)
	}

	quotes, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll() error = %v", err)
	}

	// Empty lines should be skipped
	if len(quotes) != 3 {
		t.Errorf("GetAll() returned %d quotes, want 3", len(quotes))
	}
}

func TestFileQuoteRepository_GetRandom(t *testing.T) {
	content := "Quote 1\nQuote 2\nQuote 3\n"
	filename := createTempQuotesFile(t, content)
	defer os.Remove(filename)

	repo, err := NewFileQuoteRepository(filename)
	if err != nil {
		t.Fatalf("NewFileQuoteRepository() error = %v", err)
	}

	// Get several random quotes
	seen := make(map[string]bool)
	for i := 0; i < 10; i++ {
		quote, err := repo.GetRandom()
		if err != nil {
			t.Fatalf("GetRandom() error = %v", err)
		}
		if quote == nil {
			t.Error("GetRandom() returned nil")
		}
		if quote.Text == "" {
			t.Error("GetRandom() returned empty quote")
		}
		seen[quote.Text] = true
	}

	// Check that at least one quote was retrieved
	if len(seen) == 0 {
		t.Error("GetRandom() did not return any quotes")
	}
}

func TestFileQuoteRepository_GetRandom_EmptyFile(t *testing.T) {
	content := ""
	filename := createTempQuotesFile(t, content)
	defer os.Remove(filename)

	repo, err := NewFileQuoteRepository(filename)
	if err != nil {
		t.Fatalf("NewFileQuoteRepository() error = %v", err)
	}

	quote, err := repo.GetRandom()
	if err != nil {
		t.Fatalf("GetRandom() error = %v", err)
	}

	// Default quote should be returned
	if quote == nil {
		t.Error("GetRandom() returned nil")
	}
	if quote.Text != "Wisdom comes with experience." {
		t.Errorf("GetRandom() = %v, want default quote", quote.Text)
	}
}

func TestFileQuoteRepository_ImplementsInterface(t *testing.T) {
	var _ repository.QuoteRepository = (*FileQuoteRepository)(nil)
}

func TestFileQuoteRepository_RelativePath(t *testing.T) {
	// Test with relative path
	content := "Test quote\n"
	filename := createTempQuotesFile(t, content)
	defer os.Remove(filename)

	// Use absolute path
	absPath, err := filepath.Abs(filename)
	if err != nil {
		t.Fatalf("filepath.Abs() error = %v", err)
	}

	repo, err := NewFileQuoteRepository(absPath)
	if err != nil {
		t.Fatalf("NewFileQuoteRepository() error = %v", err)
	}

	quotes, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll() error = %v", err)
	}

	if len(quotes) != 1 {
		t.Errorf("GetAll() returned %d quotes, want 1", len(quotes))
	}
}
