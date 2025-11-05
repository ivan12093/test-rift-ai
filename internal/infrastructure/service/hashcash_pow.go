package service

import (
	cryptorand "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"word-of-wisdom/internal/domain/entity"
)

type HashCashPOW struct {
	difficulty int
}

func NewHashCashPOW(difficulty int) *HashCashPOW {
	return &HashCashPOW{
		difficulty: difficulty,
	}
}

func (h *HashCashPOW) GenerateChallenge() (*entity.Challenge, error) {
	data := make([]byte, 16)
	_, err := cryptorand.Read(data)
	if err != nil {
		return nil, fmt.Errorf("failed to generate challenge: %w", err)
	}

	challengeValue := hex.EncodeToString(data)
	return entity.NewChallenge(challengeValue, h.difficulty), nil
}

func (h *HashCashPOW) Verify(challenge *entity.Challenge, solution string) (*entity.POWResult, error) {
	hash := sha256.Sum256([]byte(challenge.Value + solution))
	hashStr := hex.EncodeToString(hash[:])

	// Convert difficulty bits to number of hex characters (4 bits = 1 hex character)
	requiredZeros := challenge.Difficulty / 4
	if challenge.Difficulty%4 != 0 {
		requiredZeros++
	}

	prefix := strings.Repeat("0", requiredZeros)
	valid := strings.HasPrefix(hashStr, prefix)

	return entity.NewPOWResult(valid, nil), nil
}

func (h *HashCashPOW) Solve(challenge *entity.Challenge) (string, error) {
	// Convert difficulty bits to number of hex characters (4 bits = 1 hex character)
	requiredZeros := challenge.Difficulty / 4
	if challenge.Difficulty%4 != 0 {
		requiredZeros++
	}
	prefix := strings.Repeat("0", requiredZeros)

	nonce := 0
	for {
		solution := fmt.Sprintf("%d", nonce)
		hash := sha256.Sum256([]byte(challenge.Value + solution))
		hashStr := hex.EncodeToString(hash[:])

		if strings.HasPrefix(hashStr, prefix) {
			return solution, nil
		}

		nonce++
	}
}
