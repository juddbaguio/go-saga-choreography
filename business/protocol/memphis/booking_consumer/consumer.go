package booking_consumer

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/juddbaguio/go-saga-choreography/business/app_const"
	"github.com/juddbaguio/go-saga-choreography/business/app_const/events"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
	"github.com/juddbaguio/go-saga-choreography/business/domain/usecases/booking_usecase"
	"github.com/memphisdev/memphis.go"
	"go.uber.org/zap"
)

type Server struct {
	logger         *zap.Logger
	memphisConn    *memphis.Conn
	bookingService *booking_usecase.Container
}

func New(logger *zap.Logger, memphisConn *memphis.Conn, bookingService *booking_usecase.Container) *Server {
	return &Server{
		logger:         logger,
		memphisConn:    memphisConn,
		bookingService: bookingService,
	}
}

func (mc *Server) Run() error {
	var eventHandlerMap map[events.Event]func(booking domain.Booking) error = make(map[events.Event]func(booking domain.Booking) error)
	eventHandlerMap[events.UNBLOCKED_CINEMA_SEATS] = mc.bookingService.CancelBooking
	eventHandlerMap[events.PAYMENT_SUCCESS] = mc.bookingService.ConfirmBooking

	handle := &handlerWrapper{
		logger:          mc.logger,
		eventHandlerMap: eventHandlerMap,
	}

	booking, err := mc.memphisConn.CreateConsumer(app_const.STATION_NAME, app_const.BOOKING_CONSUMER, memphis.PullInterval(10*time.Second))
	if err != nil {
		mc.logger.Error(err.Error())
		return err
	}

	defer booking.Destroy()
	booking.Consume(handle.HandleBookingConsumption)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigCh

	mc.logger.Info(sig.String())
	mc.logger.Warn("[CONSUMER SERVER] Starting Graceful Shutdown")
	defer mc.logger.Info("[CONSUMER SERVER] shutdown complete")

	return nil
}
