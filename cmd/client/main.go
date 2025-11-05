package main

import (
	"log"
	"word-of-wisdom/config"
	"word-of-wisdom/internal/application/usecase"
	"word-of-wisdom/internal/infrastructure/service"
	"word-of-wisdom/internal/presentation/client"
)

func main() {
	cfg := config.LoadClientConfig()

	powService := service.NewHashCashPOW(0)

	solvePOWUC := usecase.NewSolvePOWUseCase(powService)

	tcpClient := client.NewTCPClient(cfg, solvePOWUC)

	if err := tcpClient.Connect(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
