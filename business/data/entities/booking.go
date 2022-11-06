package entities

import "github.com/juddbaguio/go-saga-choreography/business/app_const"

type Booking struct {
	ID       int                     `gorm:"primaryKey;column:booking_id"`
	CinemaID int                     `gorm:"column:cinema_id"`
	Date     string                  `gorm:"type:date;column:booking_date"`
	TimeSlot string                  `gorm:"column:time_slot"`
	Status   app_const.BookingStatus `gorm:"type:varchar(50);column:booking_status"`
}

func (b *Booking) TableName() string {
	return "Booking"
}
