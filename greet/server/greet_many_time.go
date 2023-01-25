package main

import (
	"fmt"
	pb "gogrpc/greet/proto"
	"log"
)

/*
Here we are going to implement gRPC functions

to find the function signature

go to greet/proto/greet_grpc.pb.go
> search for "GreetServiceServer"
	> 	Greet(context.Context, *GreetRequest) (*GreetResponse, error)


*/

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("[server streamer] GreetManyTimes called with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("hello %s number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
