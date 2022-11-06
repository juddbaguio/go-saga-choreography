package http

import (
	"log"
	"os"
)

type Server interface {
	ListenAndServe() error
}

func Start(server Server) {
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
