package main

import (
	"log"

	appPkg "github.com/user/gen-golang-fiber-by-cursor/internal/app"
	httpHandler "github.com/user/gen-golang-fiber-by-cursor/internal/handler/http"
	helloUsecase "github.com/user/gen-golang-fiber-by-cursor/internal/usecase/hello"
)

func main() {
	server := appPkg.NewServer()

	// DI
	uc := helloUsecase.New()
	h := httpHandler.NewHelloHandler(uc)
	h.Register(server.App())

	if err := server.Start(":8080"); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}
}
