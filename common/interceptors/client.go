package interceptors

import (
	"context"
	"github.com/pkg/errors"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ClientErrorInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			grpcStatus, _ := status.FromError(err)
			xc := errorx.GrpcStatusToErrorX(grpcStatus)
			err = errors.WithMessage(xc, grpcStatus.Message())
		}
		return err
	}
}
