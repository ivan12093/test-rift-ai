package entity

import "testing"

func TestNewQuote(t *testing.T) {
	text := "Test quote"
	quote := NewQuote(text)

	if quote == nil {
		t.Fatal("NewQuote() returned nil")
	}
	if quote.Text != text {
		t.Errorf("NewQuote() text = %v, want %v", quote.Text, text)
	}
}

func TestQuote_Empty(t *testing.T) {
	quote := NewQuote("")
	if quote.Text != "" {
		t.Error("NewQuote() should accept empty string")
	}
}
