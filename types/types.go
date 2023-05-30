package types

type CreateShortURLRequest struct {
	OriginalURL string `json:"originalURL" binding:"required"`
}

type CreateShortURLResponse struct {
	ShortURL string `json:"shortURL"`
}

type GetLongURLRequest struct {
	ShortURL string `json:"shortURL" binding:"required"`
}

type GetLongURLResponse struct {
	LongURL string `json:"longURL"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
