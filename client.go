package main

import (
	"context"
	"fmt"
	"log"
	pb "github.com/siddhardhk24/bank_customer_proto/netxd_customer"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomer_ServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &pb.Customer{CustomerId: 11})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response)
}
