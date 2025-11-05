package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"
	"word-of-wisdom/config"
)

const (
	maxConnections  = 100
	shutdownTimeout = 30 * time.Second
)

type TCPServer struct {
	handler *ConnectionHandler
	port    string
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
	sem     chan struct{}
}

func NewTCPServer(handler *ConnectionHandler, cfg *config.Config) *TCPServer {
	ctx, cancel := context.WithCancel(context.Background())
	return &TCPServer{
		handler: handler,
		port:    cfg.ServerPort,
		ctx:     ctx,
		cancel:  cancel,
		sem:     make(chan struct{}, maxConnections),
	}
}

func (s *TCPServer) Start() error {
	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("TCP server Word of Wisdom started on port %s", s.port)
	log.Printf("Max concurrent connections: %d", maxConnections)

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			select {
			case <-s.ctx.Done():
				return
			default:
				conn, err := listener.Accept()
				if err != nil {
					select {
					case <-s.ctx.Done():
						return
					default:
						log.Printf("Error accepting connection: %v", err)
						continue
					}
				}

				select {
				case s.sem <- struct{}{}:
					s.wg.Add(1)
					go func(c net.Conn) {
						defer s.wg.Done()
						defer func() { <-s.sem }()
						s.handler.Handle(c)
					}(conn)
				case <-s.ctx.Done():
					conn.Close()
					return
				}
			}
		}
	}()

	<-s.ctx.Done()
	log.Println("Received shutdown signal, closing listener...")
	listener.Close()

	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("All connections closed gracefully")
	case <-time.After(shutdownTimeout):
		log.Printf("Shutdown timeout (%v) exceeded, some connections may still be active", shutdownTimeout)
	}

	return nil
}

func (s *TCPServer) Shutdown() {
	log.Println("Initiating graceful shutdown...")
	s.cancel()
}
