package main

import (
	"context"
	pb "gogrpc/greet/proto"
	"log"
)

// this function consumes the gRPC endpoint
func doGreet(client pb.GreetServiceClient){
	log.Println("calling grpc function")

	res, err:= client.Greet(context.Background(),&pb.GreetRequest{
		FirstName: "Sumit",  // from proto file     rpc Greet(GreetRequest) returns (GreetResponse);

	})

	if err != nil{
		log.Fatalf("error in greet call %v\n",err)
	}

	log.Printf("greeting: %s", res.Result)
}