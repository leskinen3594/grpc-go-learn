package services

import (
	context "context"
	"fmt"
)

type handshakeServer struct {
}

func NewHandshakeServer() HandshakeServer {
	return handshakeServer{}
}

func (handshakeServer) mustEmbedUnimplementedHandshakeServer() {}

func (handshakeServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("Hello, %v", req.Name)
	res := HelloResponse {
		Result: result,
	}

	return &res, nil
}