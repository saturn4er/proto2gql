package interceptors

import (
	"fmt"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestDefaultInterceptor_OnCall(t *testing.T) {
	i := InterceptorHandler{}
	i.OnCall(func(ctx *Context, req interface{}, next CallMethodInvoker, opts ...grpc.CallOption) (result interface{}, err error) {
		res, err := next(req, opts...)
		fmt.Println(res, err)
		return res, err
	})
	i.OnCall(func(ctx *Context, req interface{}, next CallMethodInvoker, opts ...grpc.CallOption) (result interface{}, err error) {
		next(req, append(opts, grpc.Trailer(&metadata.MD{}))...)
		return "My own way", nil
	})
	res, err := i.Call(nil, 123, func(ctx *Context, req interface{}, next CallMethodInvoker, opts ...grpc.CallOption) (result interface{}, err error) {
		fmt.Println(opts)
		return "Hello", nil
	})

	fmt.Println(3, res, err)
}
