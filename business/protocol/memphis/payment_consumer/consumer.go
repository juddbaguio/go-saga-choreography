package payment_consumer

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/juddbaguio/go-saga-choreography/business/app_const"
	"github.com/juddbaguio/go-saga-choreography/business/app_const/events"
	"github.com/juddbaguio/go-saga-choreography/business/domain"
	"github.com/juddbaguio/go-saga-choreography/business/domain/usecases/payment_usecase"
	"github.com/memphisdev/memphis.go"

	"go.uber.org/zap"
)

type Server struct {
	logger         *zap.Logger
	memphisConn    *memphis.Conn
	paymentService *payment_usecase.Container
}

func New(logger *zap.Logger, memphisConn *memphis.Conn, paymentService *payment_usecase.Container) *Server {
	return &Server{
		logger:         logger,
		memphisConn:    memphisConn,
		paymentService: paymentService,
	}
}

func (mc *Server) Run() error {
	var eventHandlerMap map[events.Event]func(booking domain.Booking) error = make(map[events.Event]func(booking domain.Booking) error)
	eventHandlerMap[events.BLOCKED_CINEMA_SEATS] = mc.paymentService.CreatePayment
	eventHandlerMap[events.BOOKING_CANCELLED] = mc.paymentService.RefundPayment

	handle := &handlerWrapper{
		logger:          mc.logger,
		eventHandlerMap: eventHandlerMap,
	}

	payment, err := mc.memphisConn.CreateConsumer(app_const.STATION_NAME, app_const.PAYMENT_CONSUMER, memphis.PullInterval(10*time.Second))
	if err != nil {
		mc.logger.Error(err.Error())
		return err
	}

	defer payment.Destroy()
	payment.Consume(handle.HandlePaymentConsumption)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigCh

	mc.logger.Info(sig.String())
	mc.logger.Warn("[CONSUMER SERVER] Starting Graceful Shutdown")
	defer mc.logger.Info("[CONSUMER SERVER] shutdown complete")

	return nil
}
