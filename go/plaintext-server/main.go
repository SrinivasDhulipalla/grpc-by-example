package main

import (
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

	//pb.RegisterMyService(server, implementation)

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err.Error())
	}

	err = server.Serve(listener)
	if err != nil {
		panic(err.Error())
	}
}