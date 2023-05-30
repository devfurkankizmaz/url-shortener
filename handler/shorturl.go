package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/devfurkankizmaz/url-shortener/service"
	models "github.com/devfurkankizmaz/url-shortener/types"
	"github.com/labstack/echo/v4"
)

type ShortURLHandler struct {
	service service.ShortURLServiceImpl
}

func NewShortURLHandler(service service.ShortURLServiceImpl) *ShortURLHandler {
	return &ShortURLHandler{service: service}
}

func (h *ShortURLHandler) CreateShortURL(c echo.Context) error {
	var payload *models.CreateShortURLRequest
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	short, err := h.service.CreateShortURL(ctx, payload.OriginalURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, models.CreateShortURLResponse{ShortURL: short})
}

func (h *ShortURLHandler) FetchLongURL(c echo.Context) error {
	var payload *models.GetLongURLRequest
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	long, err := h.service.GetLongURL(ctx, payload.ShortURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, models.GetLongURLResponse{LongURL: long})
}
