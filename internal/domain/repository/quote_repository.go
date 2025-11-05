package repository

import "word-of-wisdom/internal/domain/entity"

// QuoteRepository interface for working with quotes
type QuoteRepository interface {
	GetAll() ([]*entity.Quote, error)
	GetRandom() (*entity.Quote, error)
}
