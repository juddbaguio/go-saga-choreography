package booking_usecase

import (
	"github.com/juddbaguio/go-saga-choreography/business/interface/repository"
	"github.com/memphisdev/memphis.go"
)

type Container struct {
	bookingRepo repository.Booking
	producer    *memphis.Producer
}

func New(bookingRepo repository.Booking, producer *memphis.Producer) *Container {
	return &Container{
		bookingRepo: bookingRepo,
		producer:    producer,
	}
}
