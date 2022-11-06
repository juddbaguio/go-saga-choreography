package repository

import "github.com/juddbaguio/go-saga-choreography/business/domain"

type Cinema interface {
	GetCinemaList() (*[]domain.Cinema, error)
	BlockSeats(booking domain.Booking) error
	UnblockSeats(bookingId int) error
}
