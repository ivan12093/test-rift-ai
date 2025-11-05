package entity

type Challenge struct {
	Value     string
	Difficulty int
}

func NewChallenge(value string, difficulty int) *Challenge {
	return &Challenge{
		Value:     value,
		Difficulty: difficulty,
	}
}
