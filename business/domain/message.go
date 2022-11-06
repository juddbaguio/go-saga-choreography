package domain

import "github.com/juddbaguio/go-saga-choreography/business/app_const/events"

type BookingTxMessage struct {
	Event   events.Event `json:"event"`
	Payload Booking      `json:"booking"`
}
