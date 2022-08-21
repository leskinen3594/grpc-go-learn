package main

import (
	"client/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	handshakeClient := services.NewHandshakeClient(cc)
	handshakeService := services.NewHandshakeService(handshakeClient)

	err = handshakeService.Hello("Peter")
	if err != nil {
		log.Fatal(err)
	}
}