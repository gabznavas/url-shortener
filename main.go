package main

import (
	"context"
	"net/http"
	cmd "urlshortener/cmd/url"
	urlRepositories "urlshortener/infra/repositories/url"
	createURLHandlers "urlshortener/presenters/handlers/url/create"
	createURLUsecases "urlshortener/usecases/url/create"

	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()

	urlRepository := urlRepositories.NewURLRepository()
	createURLUsecase := createURLUsecases.NewCreateURLUsecase(urlRepository)
	createURLHandler := createURLHandlers.NewCreateURLHandler(createURLUsecase)

	initCmd := cmd.NewInitCmd(urlRepository)
	initCmd.Execute(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/url", createURLHandler.CreateURL).Methods("POST")
	http.ListenAndServe(":8080", router)
}
