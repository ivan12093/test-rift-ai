package protocol

import (
	"fmt"
	"strconv"
	"strings"
)

type MessageType string

const (
	MessageChallenge MessageType = "CHALLENGE"
	MessageSolution  MessageType = "SOLUTION"
	MessageQuote     MessageType = "QUOTE"
	MessageError     MessageType = "ERROR"
)

type Message struct {
	Type      MessageType
	Challenge string
	Difficulty int
	Solution  string
	Quote     string
	Error     string
}

func ParseChallenge(data string) (*Message, error) {
	data = strings.TrimSpace(data)
	if !strings.HasPrefix(data, "CHALLENGE:") {
		return nil, fmt.Errorf("invalid challenge format")
	}

	parts := strings.Split(data, ":")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid challenge format")
	}

	msg := &Message{
		Type:      MessageChallenge,
		Challenge: parts[1],
	}
	
	var err error
	msg.Difficulty, err = parseInt(parts[2])
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func ParseSolution(data string) (*Message, error) {
	data = strings.TrimSpace(data)
	if !strings.HasPrefix(data, "SOLUTION:") {
		return nil, fmt.Errorf("invalid solution format")
	}

	parts := strings.Split(data, ":")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid solution format")
	}

	return &Message{
		Type:      MessageSolution,
		Challenge: parts[1],
		Solution:  parts[2],
	}, nil
}

func FormatChallenge(challenge string, difficulty int) string {
	return fmt.Sprintf("CHALLENGE:%s:%d\n", challenge, difficulty)
}

func FormatSolution(challenge, solution string) string {
	return fmt.Sprintf("SOLUTION:%s:%s\n", challenge, solution)
}

func FormatQuote(quote string) string {
	return fmt.Sprintf("QUOTE:%s\n", quote)
}

func FormatError(errMsg string) string {
	return fmt.Sprintf("ERROR:%s\n", errMsg)
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}
