# Customer-Waiter-Chef

A little experiment with microservices and message bus (NSQ)

# Update

There also another version of the experiment, only this time I used gRPC.

https://github.com/CJ-Jackson/customer-waiter-chef/tree/master/grpc

# Usage

Make sure that NSQ is running somewhere.

Run the following two commands first, make sure the NSQ address is correct.

```sh
$ go run cmd/chef/main.go 127.0.0.1:1450
$ go run cmd/waiter/main.go 127.0.0.1:1450
```

Than finally run the command.

```sh
$ go run cmd/customer/main.go 127.0.0.1:1450
```

# Blog Post

https://www.cj-jackson.com/blogentry/26-experiment-with-microservice-and-message-bus-nsq