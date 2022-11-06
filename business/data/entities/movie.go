package entities

type Movie struct {
	ID              int    `gorm:"primaryKey"`
	Title           string `gorm:"title"`
	StartAiringDate string `gorm:"start_airing_date"` // in UTC
	EndAiringDate   string `gorm:"end_airing_date"`   // in UTC
}
