package repositories

import (
	"context"
	"errors"
	"time"
	"urlshortener/infra/entities"

	"github.com/google/uuid"
)

var (
	ErrURLNotFound = errors.New("url not found")
)

type URLRepository interface {
	CreateURL(ctx context.Context, url string) (*entities.URL, error)
	GetURL(ctx context.Context, id string) (string, error)
	GetNextNotReadURL(ctx context.Context, url string) (*entities.URL, error)
}

type urlRepository struct {
	urls map[string]*entities.URL
}

func NewURLRepository() URLRepository {
	return &urlRepository{urls: make(map[string]*entities.URL)}
}

func (r *urlRepository) CreateURL(ctx context.Context, url string) (*entities.URL, error) {
	id := uuid.New().String()
	urlEntity := &entities.URL{
		ID:        id,
		URL:       url,
		ShortURL:  id,
		Read:      false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	r.urls[id] = urlEntity
	return urlEntity, nil
}

func (r *urlRepository) GetURL(ctx context.Context, id string) (string, error) {
	urlEntity, ok := r.urls[id]
	if !ok {
		return "", ErrURLNotFound
	}
	return urlEntity.URL, nil
}

func (r *urlRepository) GetNextNotReadURL(ctx context.Context, url string) (*entities.URL, error) {
	for _, urlEntity := range r.urls {
		if urlEntity.Read {
			continue
		}
		urlEntity.Read = true
		urlEntity.UpdatedAt = time.Now()
		return urlEntity, nil
	}
	return nil, errors.New("no next not read url found")
}
