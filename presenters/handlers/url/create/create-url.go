package handlers

import (
	"encoding/json"
	"net/http"
	dtos "urlshortener/presenters/dtos/urls"
	usecases "urlshortener/usecases/url/create"
)

type CreateURLHandler interface {
	CreateURL(w http.ResponseWriter, r *http.Request)
}

type createURLHandler struct {
	usecase usecases.CreateURLUsecase
}

func NewCreateURLHandler(usecase usecases.CreateURLUsecase) CreateURLHandler {
	return &createURLHandler{usecase: usecase}
}

func (h *createURLHandler) CreateURL(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateURLRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request body"))
		return
	}

	urlEntity, err := h.usecase.Execute(r.Context(), req.URL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := dtos.CreateURLResponse{
		ShortURL: urlEntity.ShortURL,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
