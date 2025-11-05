package usecase

import (
	"word-of-wisdom/internal/domain/entity"
)

type SolvePOWUseCase struct {
	powService POWService
}

func NewSolvePOWUseCase(powService POWService) *SolvePOWUseCase {
	return &SolvePOWUseCase{
		powService: powService,
	}
}

func (uc *SolvePOWUseCase) Execute(challengeValue string, difficulty int) (string, error) {
	challenge := entity.NewChallenge(challengeValue, difficulty)
	return uc.powService.Solve(challenge)
}
