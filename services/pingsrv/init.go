package main

import (
)
import (
	"github.com/astaxie/beego"
	"net/rpc"
	"net"
)

func main() {
	srv := new(PingService)
	port := beego.AppConfig.String("pingsrv.port")
	if port == ""{
		panic("pingsrv.port is not defined in config")
	}
	Run(srv, port)
}

func Run(srv interface{}, port string) {
	rpc.Register(srv)
	l, e := net.Listen("tcp", ":"+port)
	if e != nil {
		beego.BeeLogger.Error("Run service error. Unable to listen tcp port:", e)
	}
	beego.BeeLogger.Info("Run RPC server at port %s\n", port)
	go rpc.Accept(l)
}