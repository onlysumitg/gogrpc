package main

import (
	"context"
	pb "gogrpc/greet/proto"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*

EXCEPTION HANDLING

*/

// this function consumes the gRPC endpoint which can return an error
func doSqrt(client pb.GreetServiceClient, n int32) {
	log.Println("calling grpc function")

	res, err := client.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n, // from proto file     rpc Greet(GreetRequest) returns (GreetResponse);

	})

	if err != nil {

		e, isGrpcError := status.FromError(err)

		if isGrpcError {
			log.Printf("gRPC Error Message from server: %s\n", e.Message())
			log.Printf("gRPC Error Code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println(" have sent a negative number?")

			}

		} else {
			log.Fatalf("NON-gRPC error in SQRT call %v\n", err)
		}

	} else {

		log.Printf("SQRT result: %d", res.Result)
	}
}
