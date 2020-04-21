package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"

	pb "github.com/AkezhanOb1/auth/api"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50052", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewCompanyServicesClient(cc)

	//request := &pb.GenerateTokenRequest{
	//	Credentials: &pb.ClientCredentials{
	//		Email:    "esbolatov@gmail.com",
	//		Password: "Yfehsp2203",
	//	},
	//}

	request := &pb.RetrieveTokenInformationRequest{
		AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJlc2JvbGF0b3ZAZ21haWwuY29tIiwiZXhwIjoxNTg3NTA2Njg1LCJpYXQiOjE1ODc0OTIyODUsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6OTAwMS90b2tlbiJ9.UDVkYEK1aDCGrgWJ4VBN4bXOd_5v6kul1pEYqxpjxpA",
	}

	resp, _ := client.RetrieveTokenInformation(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp)
}