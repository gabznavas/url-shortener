package cmd

import (
	"context"
	"errors"
	"log"
	repositories "urlshortener/infra/repositories/url"
)

type InitCmd interface {
	Execute(ctx context.Context)
}

type initCmd struct {
	urlRepository repositories.URLRepository
}

func NewInitCmd(urlRepository repositories.URLRepository) InitCmd {
	return &initCmd{urlRepository: urlRepository}
}

func (initCmd *initCmd) Execute(ctx context.Context) {
	log.Println("Iniciando comando de inicialização")
	totalCaracteresAlfabeticos := 26
	urlBase := []byte("AAAAA") // existe 26^5 =  11.881.376 URLs
	for i := 0; i < totalCaracteresAlfabeticos; i++ {
		for j := 0; j < 5; j++ {
			urlBase[j] = byte(rune('A' + i))
			_, err := initCmd.urlRepository.GetURL(ctx, string(urlBase))
			if err != nil {
				if errors.Is(err, repositories.ErrURLNotFound) {
					_, err = initCmd.urlRepository.CreateURL(ctx, string(urlBase))
					if err != nil {
						log.Fatalf("failed to create url: %v", err)
					}
				} else {
					continue
				}
			}
		}
	}
	log.Println("Comando de inicialização finalizado")
}
