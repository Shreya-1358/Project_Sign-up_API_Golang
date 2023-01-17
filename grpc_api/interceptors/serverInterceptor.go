package interceptors

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"grpc/config"
)

func ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	reqStr := fmt.Sprintf("%v", req)
	logger := config.GetLogger()
	logger.Info(info.FullMethod, zap.String("request", reqStr))
	h, err := handler(ctx, req)
	return h, err

}
