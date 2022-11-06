package domain

import "github.com/juddbaguio/go-saga-choreography/business/app_const"

type Payment struct {
	RefID     int                     `json:"reference_id"`
	BookingID int                     `json:"booking_id"`
	Amount    float64                 `json:"amount"`
	Status    app_const.PaymentStatus `json:"status"`
	CreatedAt string                  `json:"created_at"`
	UpdatedAt string                  `json:"updated_at"`
}
