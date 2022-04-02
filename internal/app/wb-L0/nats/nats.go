package nats

import (
	"bytes"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"time"
	"wb-L0/internal/app/wb-L0/channels"
	"wb-L0/internal/app/wb-L0/config"
	"wb-L0/internal/app/wb-L0/logger"
	"wb-L0/internal/app/wb-L0/modelValidation"
	"wb-L0/internal/app/wb-L0/postgres"
	"wb-L0/internal/app/wb-L0/storage"
)

var (
	Sc stan.Conn
)

func init() {
	go func() {
		Sc = NewConn()
	}()
}

func NewConn() stan.Conn {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			logger.Log.Info("Trying to connect to nats server (∞)")

			Sc, err := stan.Connect(config.Config.Nats.ServerID,
				config.Config.Nats.ClientID,
				stan.NatsURL(config.Config.Nats.NatsUrl),
				stan.Pings(1, 5),
				stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
					logger.Log.Error("Connection lost, reason: ", reason)
					Sc.Close()
					Sc = NewConn()
				}))
			if err != nil {
				continue
			}
			logger.Log.Info("Successfuly connected to nats server")
			channels.ConnectedToNats <- true
			return Sc
		}
	}
}

func Subscribe() {
	for {
		select {
		case <-channels.ConnectedToNats:
			logger.Log.Info("Listening to nats channel")

			// Subscribe with manual ack mode, and set AckWait to 60 seconds
			aw, _ := time.ParseDuration("60s")
			_, err := Sc.Subscribe("test", func(msg *stan.Msg) {
				logger.Log.Info("Messasge recieved")
				msg.Ack() // Manual ACK

				// Unmarshal JSON that represents the Model data
				d := json.NewDecoder(bytes.NewReader(msg.Data))
				err := d.Decode(storage.Model)
				if err != nil {
					logger.Log.Error(err.Error())
					return
				}
				if d.More() {
					logger.Log.Error("Extraneous data after JSON object")
					return
				}

				if modelValidation.Validate(storage.Model) {
					return
				}

				postgres.AddToDb()

			}, stan.DurableName("durableID"),
				stan.MaxInflight(25),
				stan.SetManualAckMode(),
				stan.AckWait(aw),
			)
			if err != nil {
				logger.Log.Error(err.Error())
			}
		}
	}
}
