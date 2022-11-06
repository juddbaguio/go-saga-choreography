package app_const

type BookingStatus string

const (
	BOOKING_CREATED   BookingStatus = "CREATED"
	BOOKING_CONFIRMED BookingStatus = "CONFIRMED"
	BOOKING_CANCELLED BookingStatus = "CANCELLED"
)
