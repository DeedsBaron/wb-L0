package main

import (
	"bufio"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"wb-L0/internal/app/wb-L0/config"
)

func main() {
	sp, err := stan.Connect(config.Config.Nats.ServerID, "pub", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		fmt.Println(err.Error())
	}

	var message []byte
	f, err := os.Open("nats-streaming-publish/" + os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		message = append(message, s.Bytes()...)
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	sp.Publish("test", message) // does not return until an ack has been received from NATS Streaming
	//Close connection
	sp.Close()
}
