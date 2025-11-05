package usecase

import (
	"word-of-wisdom/internal/domain/entity"
)

type VerifyPOWUseCase struct {
	powService POWService
}

func NewVerifyPOWUseCase(powService POWService) *VerifyPOWUseCase {
	return &VerifyPOWUseCase{
		powService: powService,
	}
}

func (uc *VerifyPOWUseCase) Execute(challengeValue string, difficulty int, solution string) (*entity.POWResult, error) {
	challenge := entity.NewChallenge(challengeValue, difficulty)
	return uc.powService.Verify(challenge, solution)
}
