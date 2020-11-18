package interceptorserver

import (
	"context"
	"fmt"

	"bitbucket.org/celsa/interceptor/rpc/interceptor"
)

type RandomInterceptor struct{}

func (itc *RandomInterceptor) Intercept(ctx context.Context, msg *interceptor.Message) (*interceptor.Response, error) {
	fmt.Println(msg)
	return &interceptor.Response{
		Msg: "OK",
	}, nil
}
