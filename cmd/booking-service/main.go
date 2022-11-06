package main

import (
	"log"
	"os"

	"github.com/juddbaguio/go-saga-choreography/business/app_const"
	"github.com/juddbaguio/go-saga-choreography/business/data/repositories/booking"
	"github.com/juddbaguio/go-saga-choreography/business/domain/usecases/booking_usecase"
	"github.com/juddbaguio/go-saga-choreography/business/protocol/http"
	"github.com/juddbaguio/go-saga-choreography/business/protocol/http/http_booking"
	"github.com/juddbaguio/go-saga-choreography/business/protocol/memphis/booking_consumer"
	"github.com/juddbaguio/go-saga-choreography/foundation/logger"
	"github.com/juddbaguio/go-saga-choreography/infrastructure/database"
	memphis_client "github.com/juddbaguio/go-saga-choreography/infrastructure/memphis"
)

func main() {
	zap := logger.New()
	defer zap.Sync()

	db, err := database.ConnectPsqlDB()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if err := database.MigrateBooking(db); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	memphisConn, err := memphis_client.New()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer memphisConn.Close()

	producer, err := memphisConn.CreateProducer(app_const.STATION_NAME, app_const.BOOKING_PRODUCER)
	if err != nil {
		zap.Error(err.Error())
		return
	}

	defer producer.Destroy()

	bookingRepo := booking.NewRepo(db)
	bookingUC := booking_usecase.New(bookingRepo, producer)

	consumerServer := booking_consumer.New(zap, memphisConn, bookingUC)
	go consumerServer.Run()

	httpBookingServer := http_booking.NewHTTP(zap, bookingUC)
	http.Start(httpBookingServer)
}
