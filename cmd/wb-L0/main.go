package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"runtime"
	"time"
	"wb-L0/internal/app/wb-L0/config"
)

func main() {
	sc, err := stan.Connect(config.Config.Nats.ServerID, config.Config.Nats.ClientID, stan.NatsURL(config.Config.Nats.NatsUrl))
	if err != nil {
		log.Fatal(err.Error())
	}
	// Subscribe with manual ack mode, and set AckWait to 60 seconds
	aw, _ := time.ParseDuration("60s")
	sc.Subscribe("test", func(msg *stan.Msg) {
		msg.Ack() // Manual ACK
		//order := pb.Order{}
		//// Unmarshal JSON that represents the Order data
		//err := json.Unmarshal(msg.Data, &order)
		//if err != nil {
		//	log.Print(err)
		//	return
		//}
		// Handle the message
		log.Printf("Subscribed message from %s", string(msg.Data))

	}, stan.DurableName("durableID"),
		stan.MaxInflight(25),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)
	runtime.Goexit()
	return
}
