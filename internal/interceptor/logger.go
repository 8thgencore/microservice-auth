package interceptor

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
)

// LogInterceptorFactory logs info about requests for gRPC server.
func LogInterceptorFactory(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		res, err := handler(ctx, req)
		// Check the result and log error
		if err != nil {
			logger.Error(err.Error(), slog.String("method", info.FullMethod), slog.Any("req", req))
		}

		return res, err
	}
}
