package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"word-of-wisdom/config"
	"word-of-wisdom/internal/application/usecase"
	"word-of-wisdom/internal/presentation/protocol"
)

type TCPClient struct {
	serverAddr string
	solvePOWUC *usecase.SolvePOWUseCase
}

func NewTCPClient(cfg *config.Config, solvePOWUC *usecase.SolvePOWUseCase) *TCPClient {
	return &TCPClient{
		serverAddr: cfg.ServerAddr,
		solvePOWUC: solvePOWUC,
	}
}

func (c *TCPClient) Connect() error {
	log.Printf("Connecting to server %s...", c.serverAddr)

	conn, err := net.Dial("tcp", c.serverAddr)
	if err != nil {
		return fmt.Errorf("error connecting to server: %w", err)
	}
	defer conn.Close()

	log.Println("Connected to server")

	reader := bufio.NewReader(conn)
	challengeLine, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading challenge: %w", err)
	}

	challengeMsg, err := protocol.ParseChallenge(challengeLine)
	if err != nil {
		return fmt.Errorf("invalid challenge format: %w", err)
	}

	log.Printf("Received challenge: %s, difficulty: %d", challengeMsg.Challenge, challengeMsg.Difficulty)
	log.Println("Starting Proof of Work solution...")

	startTime := time.Now()
	solution, err := c.solvePOWUC.Execute(challengeMsg.Challenge, challengeMsg.Difficulty)
	if err != nil {
		return fmt.Errorf("error solving POW: %w", err)
	}
	duration := time.Since(startTime)

	log.Printf("Solution found in %v: %s", duration, solution)

	solutionMsg := protocol.FormatSolution(challengeMsg.Challenge, solution)
	_, err = conn.Write([]byte(solutionMsg))
	if err != nil {
		return fmt.Errorf("error sending solution: %w", err)
	}

	log.Println("Solution sent, waiting for quote...")

	response, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	response = strings.TrimSpace(response)

	if strings.HasPrefix(response, "ERROR:") {
		return fmt.Errorf("server error: %s", response[6:])
	}

	if strings.HasPrefix(response, "QUOTE:") {
		quote := response[6:]
		fmt.Println("\n" + strings.Repeat("=", 60))
		fmt.Println("WORD OF WISDOM:")
		fmt.Println(strings.Repeat("=", 60))
		fmt.Println(quote)
		fmt.Println(strings.Repeat("=", 60) + "\n")
	} else {
		log.Printf("Unexpected response from server: %s", response)
	}

	return nil
}
