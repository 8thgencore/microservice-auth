package interceptor

import (
	"context"
	"time"

	"github.com/8thgencore/microservice-auth/internal/metrics"
	"google.golang.org/grpc"
)

// MetricsInterceptor manages application metrics.
func MetricsInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	metrics.IncRequestCounter()

	timeStart := time.Now()
	res, err := handler(ctx, req)
	diffTime := time.Since(timeStart)

	if err != nil {
		metrics.IncResponseCounter("error", info.FullMethod)
		metrics.HistogramResponseTimeObserve("error", diffTime.Seconds())
	} else {
		metrics.IncResponseCounter("success", info.FullMethod)
		metrics.HistogramResponseTimeObserve("success", diffTime.Seconds())
	}

	return res, err
}
