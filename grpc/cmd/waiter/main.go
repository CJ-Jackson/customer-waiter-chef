package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/CJ-Jackson/customer-waiter-chef/grpc/protocol"
	"google.golang.org/grpc"
)

type Server struct {
	chef protocol.ChefClient
}

func (s *Server) OrderFromCustomer(ctx context.Context, request *protocol.CustomerRequest) (*protocol.WaiterReply, error) {
	fmt.Println("Waiter: ", request.Request)
	response, err := s.chef.OrderFromWaiter(ctx, &protocol.WaiterRequest{
		Request: request.Request,
	})
	if err != nil {
		log.Printf("Error when calling OrderFromWaiter: %s", err)
	}
	return &protocol.WaiterReply{
		Message: "Waiter: " + response.Message,
	}, nil
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := protocol.NewChefClient(conn)

	fmt.Println("Waiter Server")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{
		chef: c,
	}

	gprcServer := grpc.NewServer()

	protocol.RegisterWaiterServer(gprcServer, &s)

	if err := gprcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
