package repository

import (
	"bufio"
	"math/rand"
	"os"
	"time"
	"word-of-wisdom/internal/domain/entity"
	"word-of-wisdom/internal/domain/repository"
)

type FileQuoteRepository struct {
	filename string
	quotes   []*entity.Quote
}

func NewFileQuoteRepository(filename string) (repository.QuoteRepository, error) {
	repo := &FileQuoteRepository{
		filename: filename,
	}

	if err := repo.loadQuotes(); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *FileQuoteRepository) loadQuotes() error {
	file, err := os.Open(r.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			r.quotes = append(r.quotes, entity.NewQuote(line))
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (r *FileQuoteRepository) GetAll() ([]*entity.Quote, error) {
	return r.quotes, nil
}

func (r *FileQuoteRepository) GetRandom() (*entity.Quote, error) {
	if len(r.quotes) == 0 {
		return entity.NewQuote("Wisdom comes with experience."), nil
	}

	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.quotes[randSource.Intn(len(r.quotes))], nil
}
