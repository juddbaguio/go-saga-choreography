package payment

import "gorm.io/gorm"

type Repository struct {
	dbConn *gorm.DB
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{
		dbConn: db,
	}
}
