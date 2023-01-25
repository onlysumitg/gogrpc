package main

import (
	"fmt"
	pb "gogrpc/greet/proto"
	"io"
	"log"
)

/*

BI-directional STREAM API

*/

func (s *Server) GreetEveryOne(stream pb.GreetService_GreetEveryOneServer) error {
	log.Printf("[bi-direction streamer] GreetEveryOne called   \n")

	// infinite loop
	for {
		req, err := stream.Recv()

		// no more request
		if err == io.EOF {
			return nil
		}

		// any other error
		if err != nil {
			log.Printf("[GreetEveryOne] error reading stream %v \n", err)
		}

		res := fmt.Sprintf("bi-direction hello %s ||", req.FirstName)

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("[GreetEveryOne] error sending data to client stream %v \n", err)
		}
	}

}
