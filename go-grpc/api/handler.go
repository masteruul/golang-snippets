package api

import (
	"log"

	"golang.org/x/net/context"
)

//server represents the gRPC
type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Menerima pesan %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}
