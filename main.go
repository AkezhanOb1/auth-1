package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
)

func main() {
	address := "0.0.0.0:50052"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}