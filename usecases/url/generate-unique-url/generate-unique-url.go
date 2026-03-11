package generateuniqueurl

import (
	"context"
	"errors"
	repositories "urlshortener/infra/repositories/url"

	"github.com/google/uuid"
)

type GenerateUniqueURLUsecase interface {
	Execute(ctx context.Context) (string, error)
}

type generateUniqueURLUsecase struct {
	repository repositories.URLRepository
}

func NewGenerateUniqueURLUsecase(repository repositories.URLRepository) GenerateUniqueURLUsecase {
	return &generateUniqueURLUsecase{repository: repository}
}

func (u *generateUniqueURLUsecase) Execute(ctx context.Context) (string, error) {
	id := uuid.New().String()
	_, err := u.repository.GetURL(ctx, id)
	if err != nil {
		if !errors.Is(err, repositories.ErrURLNotFound) {
			return "", errors.New("failed to get url")
		}
	}
	return id, nil
}
