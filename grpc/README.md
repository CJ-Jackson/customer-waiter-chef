# Customer-Waiter-Chef (gRPC)

A little experiment with microservices and gRPC.

# Usage

Run the following two commands first

```sh
$ go run cmd/chef/main.go
$ go run cmd/waiter/main.go
```

Than finally run the command.

```sh
$ go run cmd/customer/main.go Pepperoni Pizza
```