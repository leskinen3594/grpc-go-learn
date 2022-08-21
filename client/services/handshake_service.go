package services

import (
	"context"
	"fmt"
)

type HandshakeService interface {
	Hello(name string) error
}

type handshakeService struct {
	handshakeClient HandshakeClient
}

func NewHandshakeService(handshakeClient HandshakeClient) HandshakeService {
	return handshakeService{handshakeClient}
}

func (base handshakeService) Hello(name string) error {
	request := HelloRequest {
		Name: name,
	}

	response, err := base.handshakeClient.Hello(context.Background(), &request)
	if (err != nil) {
		return err
	}

	fmt.Printf("Service: Hello\n")
	fmt.Printf("Request: %v\n", request.Name)
	fmt.Printf("Response: %v\n", response.Result)

	return nil
}