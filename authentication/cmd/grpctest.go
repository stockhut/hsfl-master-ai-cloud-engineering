package main

import (
	"context"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/authentication/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewAuthenticationClient(conn)
	request := &proto.GetAccountRequest{Name: "test"}

	res, err := client.GetAccount(context.Background(), request)
	if err != nil {
		log.Fatalf("could not call rpc: %v", err)
	}

	fmt.Println(res)
}
