package main

import (
	"wb-L0/internal/app/wb-L0/apiserver"
	"wb-L0/internal/app/wb-L0/logger"
	"wb-L0/internal/app/wb-L0/nats"
	"wb-L0/internal/app/wb-L0/postgres"
)

func main() {
	go postgres.RecoverCash()
	go nats.Subscribe()
	if err := apiserver.Server.Start(); err != nil {
		logger.Log.Fatal(err)
	}
	return
}
