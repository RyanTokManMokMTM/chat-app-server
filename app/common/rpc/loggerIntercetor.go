package rpc

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	//execute handler
	resp, err = handler(ctx, req)
	if err != nil {
		errType := errors.Cause(err)
		var e *errx.CustomError
		if errors.As(errType, &e) {
			//Custom Error
			logx.WithContext(ctx).Errorf("RPC Server Error : %+v", err)
			//Change to grpc err

			err = status.Error(codes.Code(e.GetCode()), e.GetMessage())
		} else {
			//Unknown rpc error
			logx.WithContext(ctx).Errorf("RPC Server Error : %+v", err)
		}
	}

	return
}
