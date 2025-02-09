package main

import (
	user "Golang/2025/02February/20250209/kitex/kitex_gen/user/user"
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
	svr := user.NewServer(&UserImpl{},
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
		server.WithServiceAddr(addr), // address
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
