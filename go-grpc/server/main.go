package main

import (
	"fmt"
	"log"
	"net"

	"github.com/masteruul/golang-snippets/go-grpc/api"
	"google.golang.org/grpc"
)

func main() {
	//create listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	//create a server instance
	s := api.Server{}

	grpcServer := grpc.NewServer()

	//attach the Ping service to the server
	api.RegisterPingServer(grpcServer, &s)

	//start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
