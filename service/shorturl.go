package service

import (
	"context"
	"fmt"

	"github.com/devfurkankizmaz/url-shortener/repository"
	"github.com/devfurkankizmaz/url-shortener/utils"
)

type ShortURLService struct {
	repo repository.RedisRepositoryImpl
}

type ShortURLServiceImpl interface {
	CreateShortURL(ctx context.Context, longURL string) (string, error)
	GetLongURL(ctx context.Context, shortURLID string) (string, error)
}

func NewShortURLService(repo repository.RedisRepositoryImpl) ShortURLServiceImpl {
	return &ShortURLService{
		repo: repo,
	}
}

func (s *ShortURLService) CreateShortURL(ctx context.Context, longURL string) (string, error) {
	shortURLID := utils.GenerateShortURLID()

	err := s.repo.SaveShortURL(ctx, shortURLID, longURL)
	if err != nil {
		return "", fmt.Errorf("failed to create short URL: %v", err)
	}

	return shortURLID, nil
}

func (s *ShortURLService) GetLongURL(ctx context.Context, shortURLID string) (string, error) {
	longURL, err := s.repo.GetLongURL(ctx, shortURLID)
	if err != nil {
		return "", fmt.Errorf("failed to get long URL: %v", err)
	}

	return longURL, nil
}
