package main

import (
	"log"
	"net"
	"google.golang.org/grpc"

	"github.com/AkezhanOb1/auth/services"
	pb "github.com/AkezhanOb1/auth/api"

)

func main() {
	address := "0.0.0.0:50052"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()

	pb.RegisterCompanyServicesServer(s, &services.TokenServer{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}