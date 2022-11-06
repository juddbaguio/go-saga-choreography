package main

import (
	"log"
	"os"

	"github.com/juddbaguio/go-saga-choreography/business/app_const"
	"github.com/juddbaguio/go-saga-choreography/business/data/repositories/payment"
	"github.com/juddbaguio/go-saga-choreography/business/domain/usecases/payment_usecase"
	"github.com/juddbaguio/go-saga-choreography/business/protocol/memphis/payment_consumer"
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

	if err := database.MigratePayment(db); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	memphisConn, err := memphis_client.New()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer memphisConn.Close()

	producer, err := memphisConn.CreateProducer(app_const.STATION_NAME, app_const.PAYMENT_PRODUCER)
	if err != nil {
		zap.Error(err.Error())
		return
	}

	defer producer.Destroy()

	paymentRepo := payment.NewRepo(db)
	paymentUC := payment_usecase.New(paymentRepo, producer)

	consumerServer := payment_consumer.New(zap, memphisConn, paymentUC)

	if err := consumerServer.Run(); err != nil {
		os.Exit(1)
	}

}
