package main

import (
	"Golang/2025/02February/20250210/kitex-etcd/kitex_gen/user/user"
	"Golang/2025/02February/20250210/kitex-etcd/server/Registry"
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10001")
	if err != nil {
		panic(err)
	}
	Registry.EtcdRegistry.ServiceRegister("user", "127.0.0.1:10001")
	svr := user.NewServer(&UserImpl{},
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
		server.WithServiceAddr(addr), // address
	)

	err = svr.Run()

	defer svr.Stop()
	defer Registry.EtcdRegistry.ServiceUnRegister("user")
	if err != nil {
		log.Println(err.Error())
	}
	defer fmt.Println("ok")
}
