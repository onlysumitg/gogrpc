package main

import (
	"context"
	pb "gogrpc/greet/proto"
	"io"
	"log"
)

// server stream api CONSUMER
// this function consumes the gRPC endpoint
func doGreetManyTimes(client pb.GreetServiceClient) {
	log.Println("calling grpc function : doGreetManyTimes")

	request := &pb.GreetRequest{
		FirstName: "Sumit", // from proto file     rpc Greet(GreetRequest) returns (GreetResponse);

	}

	stream, err := client.GreetManyTimes(context.Background(), request)

	if err != nil {
		log.Fatalf("error in greet call %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		// no more message
		if err == io.EOF {
			break
		}

		// any other error
		if err != nil {
			log.Fatalf("[doGreetManyTimes] error reading stream %v \n", err)
		}

		// actual response
		log.Printf("[doGreetManyTimes] response : %s\n", msg.Result)

	}
}
