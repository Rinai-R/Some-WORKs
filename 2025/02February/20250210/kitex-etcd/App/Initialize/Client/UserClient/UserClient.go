package UserClient

import (
	"Golang/2025/02February/20250210/kitex-etcd/App/Initialize/Client"
	"Golang/2025/02February/20250210/kitex-etcd/kitex_gen/user/user"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"log"
)

var UserClient user.Client
var err error

func InitUserClient() {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName("user-client"),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	err = Client.ETCD.DiscoverService("user")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(Client.ETCD.Services)
	addr := Client.ETCD.GetService("user")
	UserClient, err = user.NewClient(
		"user",
		client.WithHostPorts(addr),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user-client"}),
	)
	if err != nil {
		log.Panic("Client Init error " + err.Error())
	}
}
