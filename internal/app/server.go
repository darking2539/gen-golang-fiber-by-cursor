package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	fiber *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{})
	return &Server{fiber: app}
}

func (s *Server) App() *fiber.App { return s.fiber }

func (s *Server) Start(addr string) error {
	go func() {
		if err := s.fiber.Listen(addr); err != nil && err != http.ErrServerClosed {
			log.Fatalf("fiber.Listen: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Shutdown(ctx)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.fiber.ShutdownWithContext(ctx)
}
