package main

import (
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
)

func main()  {
	s := grpc.NewServer()

	// Run server
	listener, err := net.Listen("tcp", ":50051")
	if (err != nil) {
		log.Fatal(err)
	}

	// Add services
	services.RegisterHandshakeServer(s, services.NewHandshakeServer())

	fmt.Println("gRPC server listening on port 50051")
	err = s.Serve(listener)
	if (err != nil) {
		log.Fatal(err)
	}
}