package cinema

import (
	"github.com/juddbaguio/go-saga-choreography/business/data/entities"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
)

func (c *Repository) CreateCinema(cinema domain.Cinema) error {
	var payload *entities.Cinema = &entities.Cinema{
		FirstRow: cinema.FirstRow,
		LastRow:  cinema.LastRow,
		Columns:  cinema.Columns,
	}

	if err := c.dbConn.Create(payload).Error; err != nil {
		return err
	}

	return nil
}

func (c *Repository) BlockSeats(booking domain.Booking) error {
	var seatList []entities.CinemaSeat = []entities.CinemaSeat{}
	for _, seat := range booking.SeatList {
		seatList = append(seatList, entities.CinemaSeat{
			CinemaID:  booking.Cinema.ID,
			MovieID:   booking.Movie.ID,
			Date:      booking.Schedule.Date,
			TimeSlot:  booking.Schedule.TimeSlot,
			SeatNo:    seat.GenerateSeatNumber(),
			BookingID: booking.ID,
			IsTaken:   true,
		})
	}

	if err := c.dbConn.Create(&seatList).Error; err != nil {
		return err
	}

	return nil
}

func (c *Repository) UnblockSeats(bookingId int) error {
	if err := c.dbConn.Model(&entities.CinemaSeat{}).
		Where("booking_id = ?", bookingId).
		UpdateColumn("is_taken", false).Error; err != nil {
		return err
	}

	return nil
}
