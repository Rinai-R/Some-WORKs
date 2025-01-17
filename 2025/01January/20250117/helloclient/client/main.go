package main

import (
	service "Golang/2025/01January/20250117/helloclient/client/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
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
