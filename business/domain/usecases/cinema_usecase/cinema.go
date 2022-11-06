package cinema_usecase

import (
	"github.com/juddbaguio/go-saga-choreography/business/interface/repository"
	"github.com/memphisdev/memphis.go"
)

type Container struct {
	cinemaRepo repository.Cinema
	producer   *memphis.Producer
}

func New(cinemaRepo repository.Cinema, producer *memphis.Producer) *Container {
	return &Container{
		cinemaRepo: cinemaRepo,
		producer:   producer,
	}
}
