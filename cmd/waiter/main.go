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

	fmt.Println("Running Waiter")

	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(serviceAddress, config)

	q, _ := nsq.NewConsumer("customer_to_waiter", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message from customer, %s", message.Body)
		_ = w.Publish("waiter_to_chef", message.Body)
		return nil
	}))

	chefQ, _ := nsq.NewConsumer("chef_to_waiter", "ch", config)
	chefQ.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message from chef, %s", message.Body)
		_ = w.Publish("waiter_to_customer", message.Body)
		return nil
	}))

	err := q.ConnectToNSQD(serviceAddress)
	if err != nil {
		log.Panic("Could not connect")
	}

	err = chefQ.ConnectToNSQD(serviceAddress)
	if err != nil {
		log.Panic("Could not connect")
	}

	wg.Wait()
}
