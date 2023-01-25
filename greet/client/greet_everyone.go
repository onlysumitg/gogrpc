package main

import (
	"context"
	pb "gogrpc/greet/proto" // auto generated code
	"io"
	"log"
	"time"
)

// Bi-directional stream api CONSUMER
// this function consumes the gRPC endpoint
func doGreetEveryOne(client pb.GreetServiceClient) {
	log.Println("calling grpc bi-direction stream function : doGreetEveryOne")

	requests := []*pb.GreetRequest{

		{FirstName: "Sumit2"},

		{FirstName: "SOME2"},

		{FirstName: "ONE2"},
	}

	stream, err := client.GreetEveryOne(context.Background())

	if err != nil {
		log.Fatalf("error in GreetEveryOne call %v\n", err)
	}

	// channel
	waitc := make(chan struct{})

	// need 2 goroutine
	// one send request
	// 2nd will get the response

	// send request ----------------------------------------------------------
	go func() {

		for _, req := range requests {
			log.Printf("[GreetEveryOne] Sending request : %s\n", req.FirstName)
			stream.Send(req)
			time.Sleep(1 * time.Second)

		}

		stream.CloseSend() // to tell server client is done with requests

	}()

	// receive request ----------------------------------------------------------
	go func() {
		for {
			msg, err := stream.Recv()

			// no more message
			if err == io.EOF {
				break
			}

			// any other error
			if err != nil {
				log.Printf("[GreetEveryOne] error reading stream %v \n", err)
				break
			}

			// actual response
			log.Printf("[GreetEveryOne] response received : %s\n", msg.Result)

		}

		close(waitc)
	}()

	<-waitc // wait untill this channel is not closed

}
