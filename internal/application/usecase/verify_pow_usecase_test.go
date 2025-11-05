package usecase

import (
	"errors"
	"testing"
	"word-of-wisdom/internal/domain/entity"
)

// mockPOWServiceVerify mock for POWService with Verify methods
type mockPOWServiceVerify struct {
	result *entity.POWResult
	err    error
}

func (m *mockPOWServiceVerify) GenerateChallenge() (*entity.Challenge, error) {
	return nil, nil
}

func (m *mockPOWServiceVerify) Verify(challenge *entity.Challenge, solution string) (*entity.POWResult, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.result, nil
}

func (m *mockPOWServiceVerify) Solve(challenge *entity.Challenge) (string, error) {
	return "", nil
}

func TestVerifyPOWUseCase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		challenge string
		difficulty int
		solution  string
		result    *entity.POWResult
		err       error
		wantErr   bool
		wantValid bool
	}{
		{
			name:      "valid solution",
			challenge: "test-challenge",
			difficulty: 20,
			solution:  "12345",
			result:    entity.NewPOWResult(true, nil),
			err:       nil,
			wantErr:   false,
			wantValid: true,
		},
		{
			name:      "invalid solution",
			challenge: "test-challenge",
			difficulty: 20,
			solution:  "wrong-solution",
			result:    entity.NewPOWResult(false, nil),
			err:       nil,
			wantErr:   false,
			wantValid: false,
		},
		{
			name:      "service error",
			challenge: "test-challenge",
			difficulty: 20,
			solution:  "12345",
			result:    nil,
			err:       errors.New("verification error"),
			wantErr:   true,
			wantValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mockPOWServiceVerify{
				result: tt.result,
				err:    tt.err,
			}
			uc := NewVerifyPOWUseCase(mockService)

			result, err := uc.Execute(tt.challenge, tt.difficulty, tt.solution)

			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if result.Valid != tt.wantValid {
					t.Errorf("Execute() result.Valid = %v, want %v", result.Valid, tt.wantValid)
				}
			}
		})
	}
}
