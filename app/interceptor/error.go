package interceptor

import (
	"context"
	"runtime/debug"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// ErrorHandlerUnaryServerInterceptor returns UnaryServerInterceptor
func ErrorHandlerUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if rec := recover(); rec != nil {
				grpclog.Error(rec, string(debug.Stack()))
				err = errors.Errorf("%+v", rec)
			}
		}()
		resp, err = handler(ctx, req)
		return
	}
}
