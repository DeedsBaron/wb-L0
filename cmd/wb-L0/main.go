package main

import (
	"runtime"
	"wb-L0/internal/app/wb-L0/apiserver"
	"wb-L0/internal/app/wb-L0/logger"
	"wb-L0/internal/app/wb-L0/nats"
)

func main() {
	go nats.Subscribe()
	if err := apiserver.Server.Start(); err != nil {
		logger.Log.Fatal(err)
	}
	runtime.Goexit()
	return
}
