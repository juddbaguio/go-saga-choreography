package main

import (
	"log"
	"os"

	"github.com/juddbaguio/go-saga-choreography/business/app_const"
	"github.com/juddbaguio/go-saga-choreography/business/data/repositories/cinema"
	"github.com/juddbaguio/go-saga-choreography/business/domain/usecases/cinema_usecase"
	"github.com/juddbaguio/go-saga-choreography/business/protocol/memphis/cinema_consumer"
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

	if err := database.MigrateCinema(db); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if err := database.MigrateMovie(db); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	memphisConn, err := memphis_client.New()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer memphisConn.Close()

	producer, err := memphisConn.CreateProducer(app_const.STATION_NAME, app_const.CINEMA_PRODUCER)
	if err != nil {
		zap.Error(err.Error())
		return
	}

	defer producer.Destroy()

	cinemaRepo := cinema.NewRepo(db)
	cinemaUC := cinema_usecase.New(cinemaRepo, producer)

	consumerServer := cinema_consumer.New(zap, memphisConn, cinemaUC)

	if err := consumerServer.Run(); err != nil {
		os.Exit(1)
	}

}
