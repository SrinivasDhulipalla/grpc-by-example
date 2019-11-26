package main

import (
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	// client := pb.NewMyServiceClient(cc)
}
