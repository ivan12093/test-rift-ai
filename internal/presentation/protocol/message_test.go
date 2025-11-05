package protocol

import (
	"testing"
)

func TestParseChallenge(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		wantErr bool
		wantChallenge string
		wantDifficulty int
	}{
		{
			name:    "valid challenge",
			data:    "CHALLENGE:abc123def456:20\n",
			wantErr: false,
			wantChallenge: "abc123def456",
			wantDifficulty: 20,
		},
		{
			name:    "invalid prefix",
			data:    "WRONG:abc123:20\n",
			wantErr: true,
		},
		{
			name:    "invalid format - missing parts",
			data:    "CHALLENGE:abc123\n",
			wantErr: true,
		},
		{
			name:    "invalid format - too many parts",
			data:    "CHALLENGE:abc123:20:extra\n",
			wantErr: true,
		},
		{
			name:    "invalid difficulty",
			data:    "CHALLENGE:abc123:not-a-number\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := ParseChallenge(tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseChallenge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if msg.Challenge != tt.wantChallenge {
					t.Errorf("ParseChallenge() challenge = %v, want %v", msg.Challenge, tt.wantChallenge)
				}
				if msg.Difficulty != tt.wantDifficulty {
					t.Errorf("ParseChallenge() difficulty = %v, want %v", msg.Difficulty, tt.wantDifficulty)
				}
				if msg.Type != MessageChallenge {
					t.Errorf("ParseChallenge() type = %v, want %v", msg.Type, MessageChallenge)
				}
			}
		})
	}
}

func TestParseSolution(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		wantErr bool
		wantChallenge string
		wantSolution string
	}{
		{
			name:    "valid solution",
			data:    "SOLUTION:abc123def456:12345\n",
			wantErr: false,
			wantChallenge: "abc123def456",
			wantSolution: "12345",
		},
		{
			name:    "invalid prefix",
			data:    "WRONG:abc123:12345\n",
			wantErr: true,
		},
		{
			name:    "invalid format - missing parts",
			data:    "SOLUTION:abc123\n",
			wantErr: true,
		},
		{
			name:    "invalid format - too many parts",
			data:    "SOLUTION:abc123:12345:extra\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := ParseSolution(tt.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSolution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if msg.Challenge != tt.wantChallenge {
					t.Errorf("ParseSolution() challenge = %v, want %v", msg.Challenge, tt.wantChallenge)
				}
				if msg.Solution != tt.wantSolution {
					t.Errorf("ParseSolution() solution = %v, want %v", msg.Solution, tt.wantSolution)
				}
				if msg.Type != MessageSolution {
					t.Errorf("ParseSolution() type = %v, want %v", msg.Type, MessageSolution)
				}
			}
		})
	}
}

func TestFormatChallenge(t *testing.T) {
	challenge := "abc123def456"
	difficulty := 20
	
	formatted := FormatChallenge(challenge, difficulty)
	expected := "CHALLENGE:abc123def456:20\n"
	
	if formatted != expected {
		t.Errorf("FormatChallenge() = %v, want %v", formatted, expected)
	}

	// Check that we can parse it back
	msg, err := ParseChallenge(formatted)
	if err != nil {
		t.Fatalf("ParseChallenge() error = %v", err)
	}
	if msg.Challenge != challenge {
		t.Errorf("Round-trip challenge = %v, want %v", msg.Challenge, challenge)
	}
	if msg.Difficulty != difficulty {
		t.Errorf("Round-trip difficulty = %v, want %v", msg.Difficulty, difficulty)
	}
}

func TestFormatSolution(t *testing.T) {
	challenge := "abc123def456"
	solution := "12345"
	
	formatted := FormatSolution(challenge, solution)
	expected := "SOLUTION:abc123def456:12345\n"
	
	if formatted != expected {
		t.Errorf("FormatSolution() = %v, want %v", formatted, expected)
	}

	// Check that we can parse it back
	msg, err := ParseSolution(formatted)
	if err != nil {
		t.Fatalf("ParseSolution() error = %v", err)
	}
	if msg.Challenge != challenge {
		t.Errorf("Round-trip challenge = %v, want %v", msg.Challenge, challenge)
	}
	if msg.Solution != solution {
		t.Errorf("Round-trip solution = %v, want %v", msg.Solution, solution)
	}
}

func TestFormatQuote(t *testing.T) {
	quote := "Test wisdom quote"
	
	formatted := FormatQuote(quote)
	expected := "QUOTE:Test wisdom quote\n"
	
	if formatted != expected {
		t.Errorf("FormatQuote() = %v, want %v", formatted, expected)
	}
}

func TestFormatError(t *testing.T) {
	errMsg := "Invalid proof of work"
	
	formatted := FormatError(errMsg)
	expected := "ERROR:Invalid proof of work\n"
	
	if formatted != expected {
		t.Errorf("FormatError() = %v, want %v", formatted, expected)
	}
}
