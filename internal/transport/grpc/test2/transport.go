package test

import (
	"context"

	genTest2 "github.com/MShilenko/gen-test/gen/grpc-server/test2"
)

type transport struct {
	genTest2.UnsafeTest2ServiceServer
}

func New() *transport {
	return &transport{}
}

func (s *transport) TestHandlerV1(ctx context.Context, m *genTest2.TestHandlerV1Request) (*genTest2.TestHandlerV1Response, error) {
	return &genTest2.TestHandlerV1Response{
		Value: "hello2",
	}, nil
}
