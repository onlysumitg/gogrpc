package main

import (
	"context"
	"fmt"
	pb "gogrpc/greet/proto"
	"log"
	"reflect"
	"unsafe"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

/*
Here we are going to implement gRPC functions

to find the function signature

go to greet/proto/greet_grpc.pb.go
> search for "GreetServiceServer"
	> 	Greet(context.Context, *GreetRequest) (*GreetResponse, error)


*/

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet called with %v\n", in)

	p, _ := peer.FromContext(ctx)
	fmt.Println("CLient IP ADDRESS .....", p.Addr.String())

	mD, _ := metadata.FromIncomingContext(ctx)

	for key, value := range mD {
		fmt.Println("CLient META DATA.....", key, " :: ", value)
	}

	//printContextInternals(ctx, false)

	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}

// ------------------------------------
// to print content on a context
// ------------------------------------
func printContextInternals(ctx interface{}, inner bool) {
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				printContextInternals(reflectValue.Interface(), true)
			} else {
				fmt.Printf("field name: %+v\n", reflectField.Name)
				fmt.Printf("value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Printf("context is empty (int)\n")
	}
}
