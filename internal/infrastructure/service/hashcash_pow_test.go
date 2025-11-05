package service

import (
	"testing"
	"word-of-wisdom/internal/domain/entity"
)

func TestHashCashPOW_GenerateChallenge(t *testing.T) {
	pow := NewHashCashPOW(20)

	challenge1, err := pow.GenerateChallenge()
	if err != nil {
		t.Fatalf("GenerateChallenge() error = %v", err)
	}
	if challenge1 == nil {
		t.Fatal("GenerateChallenge() returned nil")
	}
	if challenge1.Value == "" {
		t.Error("Challenge value is empty")
	}
	if challenge1.Difficulty != 20 {
		t.Errorf("Challenge difficulty = %d, want 20", challenge1.Difficulty)
	}

	// Check that challenges are unique
	challenge2, err := pow.GenerateChallenge()
	if err != nil {
		t.Fatalf("GenerateChallenge() error = %v", err)
	}
	if challenge1.Value == challenge2.Value {
		t.Error("Generated challenges are not unique")
	}
}

func TestHashCashPOW_Verify(t *testing.T) {
	tests := []struct {
		name       string
		difficulty int
		challenge  string
		solution   string
		wantValid  bool
	}{
		{
			name:       "valid solution for difficulty 4",
			difficulty: 4,
			challenge:  "test-challenge",
			solution:   "0", // Need to find correct solution
			wantValid:  false, // Start with false, then find correct one
		},
	}

	pow := NewHashCashPOW(4)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			challenge := entity.NewChallenge(tt.challenge, tt.difficulty)
			
			// First solve the task
			solution, err := pow.Solve(challenge)
			if err != nil {
				t.Fatalf("Solve() error = %v", err)
			}

			// Verify the solution
			result, err := pow.Verify(challenge, solution)
			if err != nil {
				t.Fatalf("Verify() error = %v", err)
			}
			if !result.Valid {
				t.Error("Verify() returned invalid for correct solution")
			}

			// Verify incorrect solution
			wrongResult, err := pow.Verify(challenge, "wrong-solution")
			if err != nil {
				t.Fatalf("Verify() error = %v", err)
			}
			if wrongResult.Valid {
				t.Error("Verify() returned valid for incorrect solution")
			}
		})
	}
}

func TestHashCashPOW_Solve(t *testing.T) {
	tests := []struct {
		name       string
		difficulty int
		challenge  string
	}{
		{
			name:       "difficulty 4",
			difficulty: 4,
			challenge:  "test-challenge-1",
		},
		{
			name:       "difficulty 8",
			difficulty: 8,
			challenge:  "test-challenge-2",
		},
		{
			name:       "difficulty 12",
			difficulty: 12,
			challenge:  "test-challenge-3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pow := NewHashCashPOW(tt.difficulty)
			challenge := entity.NewChallenge(tt.challenge, tt.difficulty)

			solution, err := pow.Solve(challenge)
			if err != nil {
				t.Fatalf("Solve() error = %v", err)
			}
			if solution == "" {
				t.Error("Solve() returned empty solution")
			}

			// Check that solution is valid
			result, err := pow.Verify(challenge, solution)
			if err != nil {
				t.Fatalf("Verify() error = %v", err)
			}
			if !result.Valid {
				t.Error("Solve() produced invalid solution")
			}
		})
	}
}

func TestHashCashPOW_VerifyInvalidSolution(t *testing.T) {
	pow := NewHashCashPOW(20)
	challenge := entity.NewChallenge("test-challenge", 20)

	// Check obviously incorrect solution
	result, err := pow.Verify(challenge, "invalid-solution-12345")
	if err != nil {
		t.Fatalf("Verify() error = %v", err)
	}
	if result.Valid {
		t.Error("Verify() should return invalid for wrong solution")
	}
}

func TestHashCashPOW_Consistency(t *testing.T) {
	pow := NewHashCashPOW(16)
	challenge := entity.NewChallenge("consistent-test", 16)

	// Solve multiple times and check that solution is always valid
	for i := 0; i < 5; i++ {
		solution, err := pow.Solve(challenge)
		if err != nil {
			t.Fatalf("Solve() iteration %d error = %v", i, err)
		}

		result, err := pow.Verify(challenge, solution)
		if err != nil {
			t.Fatalf("Verify() iteration %d error = %v", i, err)
		}
		if !result.Valid {
			t.Errorf("Verify() iteration %d returned invalid", i)
		}
	}
}
