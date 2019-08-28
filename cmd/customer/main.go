package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	nsq "github.com/bitly/go-nsq"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Requires One Argument")
	}

	serviceAddress := os.Args[1]

	fmt.Println("Running Customer")

	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(serviceAddress, config)

	q, _ := nsq.NewConsumer("waiter_to_customer", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message from waiter, %s", message.Body)
		wg.Done()
		return nil
	}))

	err := q.ConnectToNSQD(serviceAddress)
	if err != nil {
		log.Panic("Could not connect")
	}

	_ = w.Publish("customer_to_waiter", []byte("Pepperoni Pizza"))

	wg.Wait()
}
