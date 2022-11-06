package cinema_usecase

import (
	"github.com/juddbaguio/go-saga-choreography/business/app_const/events"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
	"github.com/juddbaguio/go-saga-choreography/foundation/appjson"
)

// Biz logic for consuming "BOOKING_CREATED" event
func (c *Container) BlockSeats(booking domain.Booking) error {
	data, _ := appjson.EncodeJSONByte(domain.BookingTxMessage{
		Event:   events.BLOCKED_CINEMA_SEATS,
		Payload: booking,
	})

	if err := c.cinemaRepo.BlockSeats(booking); err != nil {
		return err
	}

	c.producer.Produce(data)
	return nil
}

// Biz logic for consuming "PAYMENT_REFUNDED" or "PAYMENT_FAILED" event
func (c *Container) UnblockSeats(booking domain.Booking) error {
	data, _ := appjson.EncodeJSONByte(domain.BookingTxMessage{
		Event:   events.UNBLOCKED_CINEMA_SEATS,
		Payload: booking,
	})

	if err := c.cinemaRepo.UnblockSeats(booking.ID); err != nil {
		return err
	}

	c.producer.Produce(data)
	return nil
}
