package main

import (
	"log"
	"net/http"

	"github.com/devfurkankizmaz/url-shortener/handler"
	"github.com/devfurkankizmaz/url-shortener/repository"
	"github.com/devfurkankizmaz/url-shortener/service"
	"github.com/labstack/echo/v4"
)

func main() {
	repository := repository.NewRedisRepository()
	service := service.NewShortURLService(repository)
	handler := handler.NewShortURLHandler(service)
	e := echo.New()

	e.POST("/short", handler.CreateShortURL)
	e.POST("/long", handler.FetchLongURL)

	log.Printf("Server is running on %s", "5050")
	if err := e.Start(":5050"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
