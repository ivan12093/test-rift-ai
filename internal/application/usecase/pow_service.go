package usecase

import "word-of-wisdom/internal/domain/entity"

type POWService interface {
	GenerateChallenge() (*entity.Challenge, error)
	Verify(challenge *entity.Challenge, solution string) (*entity.POWResult, error)
	Solve(challenge *entity.Challenge) (string, error)
}
