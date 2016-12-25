package models

import (
	"math"
	"net/rpc"

	"fmt"

	"github.com/astaxie/beego"
	"time"
)

type pingSrvClient struct {
	*rpc.Client
}

var PingSrvClient *pingSrvClient

func init() {
	PingSrvClient = new(pingSrvClient)
	port, err := beego.AppConfig.Int("pingsrv.port")
	if err != nil {
		panic(err)
	}
	var reconnectsCount = 0
	for {
		PingSrvClient.Client, err = rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
		if err != nil {
			reconnectsCount++
			sleepTime := time.Second * time.Duration(math.Exp(float64(reconnectsCount)))
			if sleepTime > time.Minute*5 {
				sleepTime = time.Minute*5
				reconnectsCount--
			}
			beego.BeeLogger.Info("dialing ping srv error: %v. Reconnecting in %v", err, sleepTime)
			time.Sleep(sleepTime)
		}
	}
	beego.BeeLogger.Info("Successfully connected to pingsrv")
}
