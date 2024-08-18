package test

import (
	"context"

	genTest "gitlab.com/MShilenko/gen-test/gen/http-server/test"
)

type Transport struct{}

func New() *Transport {
	return &Transport{}
}

func (s *Transport) PostTest(ctx context.Context, request genTest.PostTestRequestObject) (genTest.PostTestResponseObject, error) {
	return genTest.PostTest200JSONResponse(genTest.TestResponse{
		Id:     2,
		Result: "success",
	}), nil
}
