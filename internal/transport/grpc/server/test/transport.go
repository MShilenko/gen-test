package test

import (
	"context"

	genTest "github.com/MShilenko/gen-test/gen/grpc/server/test"
)

type transport struct {
	genTest.UnimplementedTestServiceServer
}

func New() *transport {
	return &transport{}
}

func (s *transport) TestHandlerV1(ctx context.Context, m *genTest.TestHandlerV1Request) (*genTest.TestHandlerV1Response, error) {
	return &genTest.TestHandlerV1Response{
		Value: "hello",
	}, nil
}
