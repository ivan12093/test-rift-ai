package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"word-of-wisdom/config"
	"word-of-wisdom/internal/application/usecase"
	"word-of-wisdom/internal/infrastructure/repository"
	"word-of-wisdom/internal/infrastructure/service"
	"word-of-wisdom/internal/presentation/server"
)

func main() {
	cfg := config.LoadServerConfig()

	quoteRepo, err := repository.NewFileQuoteRepository(cfg.QuotesFile)
	if err != nil {
		log.Fatalf("Error loading quotes: %v", err)
	}

	powService := service.NewHashCashPOW(cfg.Difficulty)

	generateChallengeUC := usecase.NewGenerateChallengeUseCase(powService)
	verifyPOWUC := usecase.NewVerifyPOWUseCase(powService)
	getQuoteUC := usecase.NewGetQuoteUseCase(quoteRepo)

	handler := server.NewConnectionHandler(
		generateChallengeUC,
		verifyPOWUC,
		getQuoteUC,
		cfg.Timeout,
	)

	tcpServer := server.NewTCPServer(handler, cfg)

	quotes, _ := quoteRepo.GetAll()
	log.Printf("Loaded quotes: %d", len(quotes))
	log.Printf("POW difficulty: %d bits", cfg.Difficulty)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	errChan := make(chan error, 1)
	go func() {
		if err := tcpServer.Start(); err != nil {
			errChan <- err
		}
	}()

	select {
	case sig := <-sigChan:
		log.Printf("Received signal: %v", sig)
		tcpServer.Shutdown()
	case err := <-errChan:
		log.Fatalf("Error starting server: %v", err)
	}
}
