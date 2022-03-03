package main

import (
	"crawler"
	"crawler/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	server := new(crawler.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Ошибка при запуске сервера %s", err.Error())
	}
}
