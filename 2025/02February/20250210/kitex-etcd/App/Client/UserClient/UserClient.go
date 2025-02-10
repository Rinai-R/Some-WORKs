package UserClient

import (
	"Golang/2025/02February/20250210/kitex-etcd/App/Client"
	"Golang/2025/02February/20250210/kitex-etcd/Logger"
	"Golang/2025/02February/20250210/kitex-etcd/kitex_gen/user/user"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"log"
)

var UserClient user.Client
var err error

func init() {
	err = Client.ETCD.DiscoverService("user")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(Client.ETCD.Services)
	addr := Client.ETCD.GetService("user")
	Logger.Logger.Info(addr + "123")
	UserClient, err = user.NewClient("user", client.WithHostPorts(addr))
	if err != nil {
		log.Panic("Client Init error " + err.Error())
	}
}
