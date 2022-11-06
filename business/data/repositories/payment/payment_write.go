package payment

import (
	"time"

	"github.com/juddbaguio/go-saga-choreography/business/app_const"
	"github.com/juddbaguio/go-saga-choreography/business/data/entities"
	"github.com/juddbaguio/go-saga-choreography/foundation/datetime"
)

func (p *Repository) CreatePayment(bookingId int, amount float64) error {
	dateTime := time.Now().UTC().Format(datetime.DateTime)
	if err := p.dbConn.Create(&entities.Payment{
		BookingID: bookingId,
		Amount:    amount,
		Status:    app_const.PAYMENT_PENDING,
		CreatedAt: dateTime,
		UpdatedAt: dateTime,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (p *Repository) CreditPayment(bookingId int, paymentRefId int) error {
	dateTime := time.Now().UTC().Format(datetime.DateTime)
	if err := p.dbConn.Model(&entities.Payment{}).Where("booking_id = ?", bookingId).
		Where("status = 'PENDING'").Updates(map[string]interface{}{
		"reference_id": paymentRefId,
		"status":       app_const.PAYMENT_CREDITED,
		"updated_at":   dateTime,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (p *Repository) RefundPayment(bookingId int) error {
	dateTime := time.Now().UTC().Format(datetime.DateTime)
	if err := p.dbConn.Model(&entities.Payment{}).Where("booking_id = ?", bookingId).
		Where("status = 'CREDITED'").Updates(map[string]interface{}{
		"status":     app_const.PAYMENT_REFUNDED,
		"updated_at": dateTime,
	}).Error; err != nil {
		return err
	}

	return nil
}
