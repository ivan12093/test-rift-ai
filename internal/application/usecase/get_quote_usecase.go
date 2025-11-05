package usecase

import (
	"word-of-wisdom/internal/domain/repository"
)

type GetQuoteUseCase struct {
	quoteRepo repository.QuoteRepository
}

func NewGetQuoteUseCase(quoteRepo repository.QuoteRepository) *GetQuoteUseCase {
	return &GetQuoteUseCase{
		quoteRepo: quoteRepo,
	}
}

func (uc *GetQuoteUseCase) Execute() (string, error) {
	quote, err := uc.quoteRepo.GetRandom()
	if err != nil {
		return "", err
	}
	return quote.Text, nil
}
