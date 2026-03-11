package dtos

type CreateURLRequest struct {
	URL string `json:"url"`
}

type CreateURLResponse struct {
	ShortURL string `json:"short_url"`
}
