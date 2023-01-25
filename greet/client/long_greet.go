package main

import (
	"context"
	pb "gogrpc/greet/proto" // auto generated code
	"log"
	"time"
)

// server stream api CONSUMER
// this function consumes the gRPC endpoint
func doLongGreet(client pb.GreetServiceClient) {
	log.Println("calling grpc function : goLongGreet")

	requests := []*pb.GreetRequest{

		{FirstName: "Sumit"},

		{FirstName: "SOME"},

		{FirstName: "ONE"},
	}

	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("error in LongGreet call %v\n", err)
	}

	for _, req := range requests {
		log.Printf("[LongGreet] Sending request : %s\n", req.FirstName)
		stream.Send(req)
		time.Sleep(1 * time.Second)

	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error in LongGreet response  %v\n", err)
	}

	log.Printf("[goLongGreet] response : %s\n", res.Result)

}
