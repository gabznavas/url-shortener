package entities

import "time"

type URL struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	ShortURL  string    `json:"short_url"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
