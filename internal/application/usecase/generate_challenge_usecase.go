package usecase

type GenerateChallengeUseCase struct {
	powService POWService
}

func NewGenerateChallengeUseCase(powService POWService) *GenerateChallengeUseCase {
	return &GenerateChallengeUseCase{
		powService: powService,
	}
}

func (uc *GenerateChallengeUseCase) Execute() (string, int, error) {
	challenge, err := uc.powService.GenerateChallenge()
	if err != nil {
		return "", 0, err
	}
	return challenge.Value, challenge.Difficulty, nil
}
