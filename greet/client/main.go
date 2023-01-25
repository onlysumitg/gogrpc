package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	pb "gogrpc/greet/proto"
)

var addr string = "localhost:5051"

func main() {
	// get gRPC connection

	tls := false // false for non ssl connections --> should be in sync with check tsl var in greet/server/main.go

	opts := []grpc.DialOption{}

	if tls {

		// SSL connection

		certFile := "ssh/ca.crt" // this file work as auth key --> like ssh keys

		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("[Client] faile to create credentials: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))

	} else {

		// non SSL connection. If server is using SSL, connection will fail
		//   check tsl var in greet/server/main.go

		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	}

	connection, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer connection.Close()

	// create a client (aka stub) for Greet Service
	client := pb.NewGreetServiceClient(connection)

	doGreet(client)

	// doGreetManyTimes(client)

	// doLongGreet(client)

	// doGreetEveryOne(client)

	// ERROR HANDLING
	//doSqrt(client, 16)
	//doSqrt(client, -16) // to chcek error

	// Deadline
	// doGreetWithDeadline(client, 5*time.Second) // should not error out

	// doGreetWithDeadline(client, 1*time.Second) // should error out

}
