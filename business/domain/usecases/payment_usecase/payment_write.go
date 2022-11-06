package payment_usecase

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/juddbaguio/go-saga-choreography/business/app_const/events"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
	"github.com/juddbaguio/go-saga-choreography/foundation/appjson"
)

const BASE_TICKET_PRICE float64 = 250

// Biz logic for consuming "BLOCKED_CINEMA_SEATS" event
func (c *Container) CreatePayment(booking domain.Booking) error {
	amount := BASE_TICKET_PRICE * float64(len(booking.SeatList))
	dataPayload := domain.BookingTxMessage{
		Event:   events.PAYMENT_SUCCESS,
		Payload: booking,
	}

	data, _ := appjson.EncodeJSONByte(dataPayload)

	if err := c.paymentRepo.CreatePayment(booking.ID, amount); err != nil {
		return err
	}

	refPaymentId, err := creditPaymentThirdParty(booking.Customer, amount)
	if err != nil {
		dataPayload.Event = events.PAYMENT_FAILED
		data, _ = appjson.EncodeJSONByte(dataPayload)
		c.producer.Produce(data)
		return err
	}

	if err := c.paymentRepo.CreditPayment(booking.ID, refPaymentId); err != nil {
		return err
	}

	c.producer.Produce(data)
	return nil
}

// Biz logic for consuming "BOOKING_CANCELLED" event
func (c *Container) RefundPayment(booking domain.Booking) error {
	data, _ := appjson.EncodeJSONByte(domain.BookingTxMessage{
		Event:   events.PAYMENT_REFUNDED,
		Payload: booking,
	})

	paymentDetails, err := c.paymentRepo.GetPaymentByBookingId(booking.ID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = refundPaymentThirdParty(paymentDetails.RefID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if err := c.paymentRepo.RefundPayment(booking.ID); err != nil {
		log.Println(err.Error())
		return err
	}

	c.producer.Produce(data)
	return nil
}

func creditPaymentThirdParty(customer domain.CustomerInformation, amount float64) (int, error) {
	baseSuccesLine := 0.85
	decider := rand.NormFloat64()
	if decider <= baseSuccesLine {
		return int(time.Now().Unix()), nil
	}

	return int(time.Now().Unix()), errors.New("payment failed")
}

func refundPaymentThirdParty(referencePaymentId int) error {
	baseSuccesLine := 0.85
	decider := rand.NormFloat64()
	if decider <= baseSuccesLine {
		return nil
	}

	return errors.New("refund payment failed")
}
