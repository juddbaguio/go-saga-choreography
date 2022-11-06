package booking

import (
	"github.com/juddbaguio/go-saga-choreography/business/app_const"
	"github.com/juddbaguio/go-saga-choreography/business/data/entities"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
)

func (b *Repository) CreateBooking(booking domain.Booking) (*domain.Booking, error) {
	var bookingPayload *entities.Booking = &entities.Booking{
		CinemaID: booking.Cinema.ID,
		Date:     booking.Schedule.Date,
		TimeSlot: booking.Schedule.TimeSlot,
		Status:   app_const.BOOKING_CREATED,
	}

	if err := b.dbConn.Create(bookingPayload).Error; err != nil {
		return nil, err
	}

	booking.ID = bookingPayload.ID
	return &booking, nil
}

func (b *Repository) UpdateBookingStatus(bookingId int, status app_const.BookingStatus) error {
	if err := b.dbConn.Model(&entities.Booking{}).
		Where("booking_id = ?", bookingId).
		UpdateColumn("booking_status", status).Error; err != nil {
		return err
	}

	return nil
}

func (b *Repository) UpdateBookingSchedule(bookingId int, schedule domain.BookingSchedule) error {
	var bookingSchedulePayload map[string]interface{} = map[string]interface{}{
		"booking_date": schedule.Date,
		"time_slot":    schedule.TimeSlot,
	}

	if err := b.dbConn.Model(&entities.Booking{}).
		Where("booking_id = ?", bookingId).
		UpdateColumns(bookingSchedulePayload).Error; err != nil {
		return err
	}

	return nil
}
