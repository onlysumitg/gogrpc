package main

import (
	"context"
	"fmt"
	pb "gogrpc/greet/proto"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*

EXCEPTION HANDLING


send error back to client over grpc


*/

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt called with %v\n", in)

	number := in.Number

	// raise error is input number is negative
	if number < 0 {

		// status from 	"google.golang.org/grpc/status"
		// status.Errorf  need two values :: 1. Error code 2. Message
		// Error code comes fromm google.golang.org/grpc/codes
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("-ve number now allowed %d", number))
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
