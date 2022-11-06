package domain

type Movie struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	StartAiringDate string `json:"start_airing_date"` // in UTC
	EndAiringDate   string `json:"end_airing_date"`   // in UTC
}
