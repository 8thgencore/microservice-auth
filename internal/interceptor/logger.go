package interceptor

import (
	"context"

	"github.com/8thgencore/microservice-common/pkg/logger"
	"google.golang.org/grpc"
)

// LogInterceptor logs info about requests for gRPC server.
func LogInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	res, err := handler(ctx, req)
	// Check the result and log error
	if err != nil {
		logger.Error(err.Error(), logger.String("method", info.FullMethod), logger.Any("req", req))
	}

	return res, err
}
