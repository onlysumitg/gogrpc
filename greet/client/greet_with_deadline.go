package main

import (
	"context"
	pb "gogrpc/greet/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// this function consumes the gRPC endpoint
func doGreetWithDeadline(client pb.GreetServiceClient, timeout time.Duration) {
	log.Println("calling do Greet With Deadline  ")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel() // this will be called if dealine excedded

	res, err := client.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: "Sumit3", // from proto file     rpc Greet(GreetRequest) returns (GreetResponse);

	})

	if err != nil {
		e, isGrpcError := status.FromError(err)

		if isGrpcError {
			log.Printf("gRPC Error Message from server: %s\n", e.Message())
			log.Printf("gRPC Error Code from server: %s\n", e.Code())

			if e.Code() == codes.Canceled {
				log.Println(" TIMED OUT")

			}

		} else {
			log.Fatalf("NON-gRPC error in SQRT call %v\n", err)
		}
	} else {
		log.Printf("greeting with deadline completed before deadline: %s", res.Result)

	}

}
