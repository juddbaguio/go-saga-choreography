package payment

import (
	"github.com/juddbaguio/go-saga-choreography/business/data/entities"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
)

func (p *Repository) GetPaymentByBookingId(bookingId int) (*domain.Payment, error) {
	var queryPayment *entities.Payment = &entities.Payment{}

	if err := p.dbConn.Where("booking_id = ?", bookingId).First(queryPayment).Error; err != nil {
		return nil, err
	}

	return &domain.Payment{
		RefID:     queryPayment.RefID,
		BookingID: bookingId,
		Amount:    queryPayment.Amount,
		Status:    queryPayment.Status,
		CreatedAt: queryPayment.CreatedAt,
		UpdatedAt: queryPayment.UpdatedAt,
	}, nil
}
