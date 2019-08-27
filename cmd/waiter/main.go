package main

import (
	"fmt"
	"log"
	"sync"

	nsq "github.com/bitly/go-nsq"
)

func main() {
	fmt.Println("Running Waiter")

	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("cjserver.lan:4150", config)

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

	err := q.ConnectToNSQD("cjserver.lan:4150")
	if err != nil {
		log.Panic("Could not connect")
	}

	err = chefQ.ConnectToNSQD("cjserver.lan:4150")
	if err != nil {
		log.Panic("Could not connect")
	}

	wg.Wait()
}
