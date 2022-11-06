package app_const

type PaymentStatus string

const (
	PAYMENT_CREDITED PaymentStatus = "CREDITED"
	PAYMENT_PENDING  PaymentStatus = "PENDING"
	PAYMENT_REFUNDED PaymentStatus = "REFUNDED"
)
