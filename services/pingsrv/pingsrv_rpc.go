package main

import (
	"retargetapp/libs/logger"
	"time"
)

type PingService bool

func (ts *PingService) Ping(requestTime *time.Time, duration *time.Duration) error {
	t := time.Now()
	d := t.Sub(*requestTime)

	logger.Info.Printf("Request latency: %s \n", d.String())
	*duration = d
	*requestTime = time.Now()
	return nil
}
