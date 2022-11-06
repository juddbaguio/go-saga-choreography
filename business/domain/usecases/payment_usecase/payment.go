package payment_usecase

import (
	"github.com/juddbaguio/go-saga-choreography/business/interface/repository"
	"github.com/memphisdev/memphis.go"
)

type Container struct {
	paymentRepo repository.Payment
	producer    *memphis.Producer
}

func New(paymentRepo repository.Payment, producer *memphis.Producer) *Container {
	return &Container{
		paymentRepo: paymentRepo,
		producer:    producer,
	}
}
