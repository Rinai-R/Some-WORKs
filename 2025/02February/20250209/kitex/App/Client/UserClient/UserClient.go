package UserClient

import (
	"Golang/2025/02February/20250209/kitex/kitex_gen/user/user"
	"github.com/cloudwego/kitex/client"
	"log"
)

var UserClient user.Client
var err error

func init() {
	UserClient, err = user.NewClient("user", client.WithHostPorts("127.0.0.1:10001"))
	if err != nil {
		log.Panic("Client Init error " + err.Error())
	}
}
