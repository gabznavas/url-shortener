package usecases

import (
	"context"
	"errors"
	"urlshortener/infra/entities"
	repositories "urlshortener/infra/repositories/url"
)

type CreateURLUsecase interface {
	Execute(ctx context.Context, url string) (*entities.URL, error)
}

type createURLUsecase struct {
	repository repositories.URLRepository
}

func NewCreateURLUsecase(repository repositories.URLRepository) CreateURLUsecase {
	return &createURLUsecase{repository: repository}
}

func (u *createURLUsecase) Execute(ctx context.Context, url string) (*entities.URL, error) {
	urlEntity, err := u.repository.GetNextNotReadURL(ctx, url)
	if err != nil {
		return nil, errors.New("failed to get next not read url" + err.Error())
	}
	return urlEntity, nil
}
