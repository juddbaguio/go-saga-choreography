package memphis_client

import (
	"github.com/caarlos0/env/v6"
	"github.com/juddbaguio/go-saga-choreography/business/app_const/events"
	"github.com/memphisdev/memphis.go"
)

type ConsumerConfig struct {
	Event    events.Event
	Consumer *memphis.Consumer
	Handler  memphis.ConsumeHandler
}

type ConsumerMap map[events.Event]ConsumerConfig

type config struct {
	Host     string `env:"MEMPHIS_HOST,required"`
	Username string `env:"MEMPHIS_USERNAME,required"`
	Token    string `env:"MEMPHIS_TOKEN,required"`
	Port     int    `env:"MEMPHIS_PORT" envDefault:"6666"`
}

func New() (*memphis.Conn, error) {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	c, err := memphis.Connect(cfg.Host,
		cfg.Username,
		cfg.Token, memphis.Port(cfg.Port),
		memphis.Reconnect(true),
		memphis.MaxReconnect(3))

	if err != nil {
		return nil, err
	}

	return c, nil
}
