package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/CJ-Jackson/customer-waiter-chef/grpc/protocol"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) OrderFromWaiter(ctx context.Context, request *protocol.WaiterRequest) (*protocol.ChefReply, error) {
	fmt.Println("Chef:", request.Request)
	return &protocol.ChefReply{
		Message: "Chef: " + request.Request,
	}, nil
}

func main() {
	fmt.Println("Chef Server")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	gprcServer := grpc.NewServer()

	protocol.RegisterChefServer(gprcServer, &s)

	if err := gprcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
