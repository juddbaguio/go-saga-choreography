package entities

import "github.com/juddbaguio/go-saga-choreography/business/app_const"

type Payment struct {
	ID        int                     `gorm:"primaryKey"`
	BookingID int                     `gorm:"column:booking_id"`
	RefID     int                     `gorm:"column:reference_id"`
	Amount    float64                 `gorm:"type:numeric(6,2);column:amount"`
	Status    app_const.PaymentStatus `gorm:"type:varchar(50);column:status"`
	CreatedAt string                  `gorm:"type:timestamp;column:created_at"`
	UpdatedAt string                  `gorm:"type:timestamp;column:updated_at"`
}

func (p *Payment) TableName() string {
	return "Payment"
}
