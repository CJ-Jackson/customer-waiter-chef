package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/CJ-Jackson/customer-waiter-chef/grpc/protocol"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := protocol.NewWaiterClient(conn)

	response, err := c.OrderFromCustomer(context.Background(), &protocol.CustomerRequest{
		Request: os.Args[1],
	})
	if err != nil {
		log.Printf("Error when calling OrderFromCustomer: %s", err)
	}

	fmt.Println(response.Message)
}
