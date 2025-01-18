package main

import (
	service "Golang/2025/01January/20250117/helloclient/client/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"name":     "admin",
		"password": "admin",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	//cred, _ := credentials.NewClientTLSFromFile("2025/01January/20250117/helloclient/key/test.pem",
	//	"*.rinai.com")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	client := service.NewSayHelloClient(conn)

	getmes, err := client.SayHello(context.Background(), &service.HelloRequest{RequestName: "ggb"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getmes)

}
