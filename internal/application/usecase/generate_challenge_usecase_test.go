package usecase

import (
	"errors"
	"testing"
	"word-of-wisdom/internal/domain/entity"
)

// mockPOWService mock for POWService
type mockPOWService struct {
	challenge *entity.Challenge
	err       error
}

func (m *mockPOWService) GenerateChallenge() (*entity.Challenge, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.challenge, nil
}

func (m *mockPOWService) Verify(challenge *entity.Challenge, solution string) (*entity.POWResult, error) {
	return nil, nil
}

func (m *mockPOWService) Solve(challenge *entity.Challenge) (string, error) {
	return "", nil
}

func TestGenerateChallengeUseCase_Execute(t *testing.T) {
	tests := []struct {
		name           string
		challenge      *entity.Challenge
		err            error
		wantErr        bool
		wantDifficulty int
	}{
		{
			name: "successful challenge generation",
			challenge: entity.NewChallenge("test-challenge-123", 20),
			err:   nil,
			wantErr: false,
			wantDifficulty: 20,
		},
		{
			name:     "service error",
			challenge: nil,
			err:      errors.New("service error"),
			wantErr:  true,
			wantDifficulty: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mockPOWService{
				challenge: tt.challenge,
				err:       tt.err,
			}
			uc := NewGenerateChallengeUseCase(mockService)

			challengeValue, difficulty, err := uc.Execute()

			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if challengeValue != tt.challenge.Value {
					t.Errorf("Execute() challengeValue = %v, want %v", challengeValue, tt.challenge.Value)
				}
				if difficulty != tt.wantDifficulty {
					t.Errorf("Execute() difficulty = %v, want %v", difficulty, tt.wantDifficulty)
				}
			}
		})
	}
}
