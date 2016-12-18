package main

import (
	"time"
	"github.com/astaxie/beego/logs"
)

type PingService bool

func (ts *PingService) Ping(requestTime *time.Time, duration *time.Duration) error {
	t := time.Now()
	d := t.Sub(*requestTime)

	logs.Info("Request latency: %s \n", d.String())
	*duration = d
	*requestTime = time.Now()
	return nil
}
