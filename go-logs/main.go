package main

import (
	"time"

	"github.com/moov-io/base/log"
)

func main() {
	logger := log.NewDefaultLogger()
	var i int
	for true {
		i++
		logger.With(log.Fields{
			"custom1": log.String("value1"),
			"custom2": log.String("value2"),
			"number":  log.Int(i),
		}).Logf("test")
		time.Sleep(5 * time.Second)
	}
}
