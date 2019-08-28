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

	fmt.Println("Running Chef")

	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(serviceAddress, config)

	q, _ := nsq.NewConsumer("waiter_to_chef", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message from waiter, %s", message.Body)
		_ = w.Publish("chef_to_waiter", message.Body)
		return nil
	}))

	err := q.ConnectToNSQD(serviceAddress)
	if err != nil {
		log.Panic("Could not connect")
	}

	wg.Wait()
}
