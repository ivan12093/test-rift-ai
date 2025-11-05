package usecase

import (
	"errors"
	"testing"
	"word-of-wisdom/internal/domain/entity"
)

// mockQuoteRepository mock for QuoteRepository
type mockQuoteRepository struct {
	quotes []*entity.Quote
	err    error
}

func (m *mockQuoteRepository) GetAll() ([]*entity.Quote, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.quotes, nil
}

func (m *mockQuoteRepository) GetRandom() (*entity.Quote, error) {
	if m.err != nil {
		return nil, m.err
	}
	if len(m.quotes) == 0 {
		return entity.NewQuote("Wisdom comes with experience."), nil
	}
	return m.quotes[0], nil
}

func TestGetQuoteUseCase_Execute(t *testing.T) {
	tests := []struct {
		name     string
		quotes   []*entity.Quote
		err      error
		wantErr  bool
		wantText string
	}{
		{
			name: "successful quote retrieval",
			quotes: []*entity.Quote{
				entity.NewQuote("Test quote 1"),
			},
			err:      nil,
			wantErr:  false,
			wantText: "Test quote 1",
		},
		{
			name:     "repository error",
			quotes:   nil,
			err:      errors.New("repository error"),
			wantErr:  true,
			wantText: "",
		},
		{
			name:     "empty quotes",
			quotes:   []*entity.Quote{},
			err:      nil,
			wantErr:  false,
			wantText: "Wisdom comes with experience.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockQuoteRepository{
				quotes: tt.quotes,
				err:    tt.err,
			}
			uc := NewGetQuoteUseCase(repo)

			quote, err := uc.Execute()

			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && quote != tt.wantText {
				t.Errorf("Execute() quote = %v, want %v", quote, tt.wantText)
			}
		})
	}
}
