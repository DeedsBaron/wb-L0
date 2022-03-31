package utils

import (
	"github.com/sirupsen/logrus"
	"time"
)

func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	i := 1
	if attempts == 0 {
		for {
			logrus.Info("Trying to connect to nats server ", i, "(", attempts, ")")
			if err = fn(); err != nil {
				time.Sleep(delay)
				continue
			}
			return nil
		}
	} else {
		for j := attempts; j > 0; {
			logrus.Info("Trying to connect to database attempt ", i, "(", attempts, ")")
			i += 1
			if err = fn(); err != nil {
				time.Sleep(delay)
				j--
				continue
			}
			return nil
		}
		return
	}
}
