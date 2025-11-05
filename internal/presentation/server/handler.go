package server

import (
	"bufio"
	"log"
	"net"
	"time"
	"word-of-wisdom/internal/application/usecase"
	"word-of-wisdom/internal/presentation/protocol"
)

type ConnectionHandler struct {
	generateChallengeUC *usecase.GenerateChallengeUseCase
	verifyPOWUC         *usecase.VerifyPOWUseCase
	getQuoteUC          *usecase.GetQuoteUseCase
	timeout             time.Duration
}

func NewConnectionHandler(
	generateChallengeUC *usecase.GenerateChallengeUseCase,
	verifyPOWUC *usecase.VerifyPOWUseCase,
	getQuoteUC *usecase.GetQuoteUseCase,
	timeoutSeconds int,
) *ConnectionHandler {
	return &ConnectionHandler{
		generateChallengeUC: generateChallengeUC,
		verifyPOWUC:         verifyPOWUC,
		getQuoteUC:          getQuoteUC,
		timeout:             time.Duration(timeoutSeconds) * time.Second,
	}
}

func (h *ConnectionHandler) Handle(conn net.Conn) {
	defer conn.Close()

	challengeValue, difficulty, err := h.generateChallengeUC.Execute()
	if err != nil {
		log.Printf("Error generating challenge: %v", err)
		return
	}

	challengeMsg := protocol.FormatChallenge(challengeValue, difficulty)
	_, err = conn.Write([]byte(challengeMsg))
	if err != nil {
		log.Printf("Error sending challenge: %v", err)
		return
	}

	log.Printf("Sent challenge %s to client %s", challengeValue, conn.RemoteAddr())

	conn.SetReadDeadline(time.Now().Add(h.timeout))

	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading response: %v", err)
		return
	}

	solutionMsg, err := protocol.ParseSolution(response)
	if err != nil {
		conn.Write([]byte(protocol.FormatError("Invalid format")))
		return
	}

	if solutionMsg.Challenge != challengeValue {
		conn.Write([]byte(protocol.FormatError("Challenge mismatch")))
		return
	}

	powResult, err := h.verifyPOWUC.Execute(challengeValue, difficulty, solutionMsg.Solution)
	if err != nil {
		conn.Write([]byte(protocol.FormatError("Verification error")))
		return
	}

	if !powResult.Valid {
		conn.Write([]byte(protocol.FormatError("Invalid proof of work")))
		log.Printf("Invalid POW from client %s", conn.RemoteAddr())
		return
	}

	quote, err := h.getQuoteUC.Execute()
	if err != nil {
		conn.Write([]byte(protocol.FormatError("Failed to get quote")))
		return
	}

	quoteMsg := protocol.FormatQuote(quote)
	_, err = conn.Write([]byte(quoteMsg))
	if err != nil {
		log.Printf("Error sending quote: %v", err)
		return
	}

	log.Printf("Sent quote to client %s", conn.RemoteAddr())
}
