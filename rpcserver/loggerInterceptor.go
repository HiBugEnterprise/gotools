package rpcserver

import (
	"context"
	"github.com/HiBugEnterprise/gotools/errorx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)

	if err != nil {
		var customErr *errorx.Error
		if errors.As(err, &customErr) { //自定义错误类型
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
			//转成grpc err
			err = status.Error(codes.Code(customErr.Code), customErr.Msg)
		} else {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
		}

	}

	return resp, err
}
