package repository

import (
	"github.com/juddbaguio/go-saga-choreography/business/app_const"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
)

type Booking interface {
	GetBookingList() (*[]domain.Booking, error)
	CreateBooking(booking domain.Booking) (*domain.Booking, error)
	UpdateBookingStatus(bookingId int, status app_const.BookingStatus) error
	UpdateBookingSchedule(bookingId int, schedule domain.BookingSchedule) error
}
