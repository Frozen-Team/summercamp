package main

import (
	"fmt"
	"log"
	"net/rpc"
	"retargetapp"
	"testing"
	"time"
	"github.com/astaxie/beego"
)

func TestPingsrv(t *testing.T) {
	port, err := beego.AppConfig.Int("pingsrv.port")
	if err != nil {
		panic(err)
	}
	client, err := rpc.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	for {
		duration := time.Duration(0)
		now := time.Now()
		err = client.Call("PingService.Ping", &now, &duration)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		resD := time.Now().Sub(now)
		fmt.Printf("Request latency: %v\t", duration)
		fmt.Printf("Response latency: %v\t", resD)
		fmt.Printf("Total reuqest time: %v\n", resD+duration)
		time.Sleep(time.Duration(time.Second))
	}
}
