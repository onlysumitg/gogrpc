package main

import (
	"context"
	pb "gogrpc/greet/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
Here we are going to implement gRPC functions

to find the function signature

go to greet/proto/greet_grpc.pb.go
> search for "GreetServiceServer"
	> 	Greet(context.Context, *GreetRequest) (*GreetResponse, error)


*/

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline called with %v\n", in)

	// wait for 3 sec and check if client dealline has excedded --> cancel the request
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Client cancelled the reqest")
			return nil, status.Errorf(codes.Canceled, "Client canceled the request")

		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Deadline -> Hello" + in.FirstName,
	}, nil
}
