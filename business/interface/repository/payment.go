package repository

import "github.com/juddbaguio/go-saga-choreography/business/domain"

type Payment interface {
	GetPaymentByBookingId(bookingId int) (*domain.Payment, error)
	CreatePayment(bookingId int, amount float64) error
	CreditPayment(bookingId int, paymentRefId int) error
	RefundPayment(bookingId int) error
}
