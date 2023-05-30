package models

import "context"

type ShortURLService interface {
	CreateShortURL(ctx context.Context, longURL string) (string, error)
	GetLongURL(ctx context.Context, shortURLID string) (string, error)
}

type RedisRepository interface {
	SaveShortURL(ctx context.Context, shortURLID string, longURL string) error
	GetLongURL(ctx context.Context, shortURLID string) (string, error)
}
