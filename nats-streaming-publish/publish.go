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
	sp, err := stan.Connect(config.Config.Nats.ServerID, "pub", stan.NatsURL(config.Config.Nats.NatsUrl))
	if err != nil {
		fmt.Println(err.Error())
	}

	f, err := os.Open("nats-streaming-publish/publish.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	//
	//sp.Publish("test", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming
	//
	//// Close connection
	sp.Close()
}
