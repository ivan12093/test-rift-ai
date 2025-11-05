package usecase

import (
	"errors"
	"testing"
	"word-of-wisdom/internal/domain/entity"
)

// mockPOWServiceSolve mock for POWService with Solve method
type mockPOWServiceSolve struct {
	solution string
	err      error
}

func (m *mockPOWServiceSolve) GenerateChallenge() (*entity.Challenge, error) {
	return nil, nil
}

func (m *mockPOWServiceSolve) Verify(challenge *entity.Challenge, solution string) (*entity.POWResult, error) {
	return nil, nil
}

func (m *mockPOWServiceSolve) Solve(challenge *entity.Challenge) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.solution, nil
}

func TestSolvePOWUseCase_Execute(t *testing.T) {
	tests := []struct {
		name       string
		challenge  string
		difficulty int
		solution   string
		err        error
		wantErr    bool
		wantSolution string
	}{
		{
			name:       "successful solve",
			challenge:  "test-challenge",
			difficulty: 20,
			solution:   "12345",
			err:        nil,
			wantErr:    false,
			wantSolution: "12345",
		},
		{
			name:       "solve error",
			challenge:  "test-challenge",
			difficulty: 20,
			solution:   "",
			err:        errors.New("solve error"),
			wantErr:    true,
			wantSolution: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mockPOWServiceSolve{
				solution: tt.solution,
				err:      tt.err,
			}
			uc := NewSolvePOWUseCase(mockService)

			solution, err := uc.Execute(tt.challenge, tt.difficulty)

			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && solution != tt.wantSolution {
				t.Errorf("Execute() solution = %v, want %v", solution, tt.wantSolution)
			}
		})
	}
}
