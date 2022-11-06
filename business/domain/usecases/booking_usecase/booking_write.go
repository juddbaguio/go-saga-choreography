package booking_usecase

import (
	"github.com/juddbaguio/go-saga-choreography/business/app_const"
	"github.com/juddbaguio/go-saga-choreography/business/app_const/events"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
	"github.com/juddbaguio/go-saga-choreography/foundation/appjson"
)

// Initial Biz Logic for Booking Transaction
func (c *Container) CreateBooking(booking domain.Booking) (*domain.Booking, error) {
	createdBooking, err := c.bookingRepo.CreateBooking(booking)
	if err != nil {
		return nil, err
	}

	data, err := appjson.EncodeJSONByte(domain.BookingTxMessage{
		Event:   events.BOOKING_CREATED,
		Payload: *createdBooking,
	})
	if err != nil {
		return nil, err
	}

	c.producer.Produce(data)
	return createdBooking, nil
}

// Biz logic for consuming "UNBLOCKED_CINEMA_SEATS" event
func (c *Container) CancelBooking(booking domain.Booking) error {
	err := c.bookingRepo.UpdateBookingStatus(booking.ID, app_const.BOOKING_CANCELLED)
	if err != nil {
		return err
	}

	return nil
}

// Biz Logic for consuming "PAYMENT_SUCCESS" event
func (c *Container) ConfirmBooking(booking domain.Booking) error {
	err := c.bookingRepo.UpdateBookingStatus(booking.ID, app_const.BOOKING_CONFIRMED)
	if err != nil {
		return err
	}

	return nil
}
