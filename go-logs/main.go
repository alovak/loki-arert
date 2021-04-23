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
		l := logger.With(log.Fields{
			"service": log.String("paygate"),
			"number":  log.String(fmt.Sprintf("%d", i)),
		})

		l.Logf("test")

		if i%6 == 0 {
			l.LogErrorf("failed to do something: %d", i)
		}

		time.Sleep(5 * time.Second)
	}
}
