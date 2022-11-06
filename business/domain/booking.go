package domain

import "github.com/juddbaguio/go-saga-choreography/business/app_const"

type Booking struct {
	ID       int                     `json:"id"`
	Cinema   Cinema                  `json:"cinema"`
	Movie    Movie                   `json:"movie"`
	Customer CustomerInformation     `json:"customer"`
	SeatList []Seat                  `json:"seat"`
	Schedule BookingSchedule         `json:"schedule"`
	Status   app_const.BookingStatus `json:"booking_status"`
}

type BookingSchedule struct {
	Date     string `json:"date"`      // in UTC
	TimeSlot string `json:"time_slot"` // in 24-hour format
}

type CustomerInformation struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
