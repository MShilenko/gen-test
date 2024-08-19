package test2

import (
	"context"

	gen2Test "github.com/MShilenko/gen-test/gen/http/server/test2"
)

type Transport struct{}

func New() *Transport {
	return &Transport{}
}

func (s *Transport) PostTest2(ctx context.Context, request gen2Test.PostTest2RequestObject) (gen2Test.PostTest2ResponseObject, error) {
	return gen2Test.PostTest2200JSONResponse(gen2Test.Test2Response{
		Id:     22,
		Result: "success2",
	}), nil
}
