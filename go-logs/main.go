package main

import (
	"fmt"
	"time"

	"github.com/moov-io/base/log"
)

func main() {
	logger := log.NewDefaultLogger()
	var i int
	for true {
		i++
		logger.With(log.Fields{
			"service": log.String("paygate"),
			"number":  log.String(fmt.Sprintf("%d", i)),
		}).Logf("test")

		if i%3 == 0 {
			logger.With(log.Fields{
				"service": log.String("paygate"),
			}).LogErrorf("failed to do somethgin")
		}

		time.Sleep(5 * time.Second)
	}
}
