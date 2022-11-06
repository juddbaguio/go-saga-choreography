package booking_consumer

import (
	"github.com/juddbaguio/go-saga-choreography/business/app_const/events"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
	"github.com/juddbaguio/go-saga-choreography/foundation/appjson"
	"github.com/memphisdev/memphis.go"
	"go.uber.org/zap"
)

type handlerWrapper struct {
	logger          *zap.Logger
	eventHandlerMap map[events.Event]func(booking domain.Booking) error
}

func (h *handlerWrapper) HandleBookingConsumption(msgs []*memphis.Msg, err error) {
	if err != nil {
		h.logger.Sugar().Errorf("Fetch failed: %v\n", err)
		return
	}

	for _, msg := range msgs {
		var txMsg domain.BookingTxMessage

		if err = appjson.DecodeJSONByte(msg.Data(), &txMsg); err != nil {
			return
		}

		if fn, ok := h.eventHandlerMap[txMsg.Event]; ok {
			if err := fn(txMsg.Payload); err != nil {
				h.logger.Sugar().Errorf("ERROR: %v\n", err)

				return
			}
		}

		msg.Ack()
	}
}
