package main

import (
	"fmt"
	pb "gogrpc/greet/proto"
	"io"
	"log"
)

/*

CLIENT STREAM API

*/

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("[client streamer] LongGreet called with \n")

	res := ""
	for {
		req, err := stream.Recv()

		// no more request
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		// any other error
		if err != nil {
			log.Fatalf("[doGreetManyTimes] error reading stream %v \n", err)
		}

		res += fmt.Sprintf("hello %s ||", req.FirstName)
	}

}
