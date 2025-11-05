package entity

import "testing"

func TestNewChallenge(t *testing.T) {
	value := "test-challenge-123"
	difficulty := 20

	challenge := NewChallenge(value, difficulty)

	if challenge == nil {
		t.Fatal("NewChallenge() returned nil")
	}
	if challenge.Value != value {
		t.Errorf("NewChallenge() value = %v, want %v", challenge.Value, value)
	}
	if challenge.Difficulty != difficulty {
		t.Errorf("NewChallenge() difficulty = %v, want %v", challenge.Difficulty, difficulty)
	}
}

func TestNewChallenge_ZeroDifficulty(t *testing.T) {
	challenge := NewChallenge("test", 0)
	if challenge.Difficulty != 0 {
		t.Errorf("NewChallenge() should accept zero difficulty")
	}
}
