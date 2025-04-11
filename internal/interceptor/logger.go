package interceptor

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
)

// LogInterceptorFactory logs info about requests for gRPC server.
func LogInterceptorFactory(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		res, err := handler(ctx, req)
		if err != nil {
			reqStr := fmt.Sprintf("%v", req)
			reqStr = maskSensitiveFields(reqStr, []string{"password", "confirm_password"})
			logger.Error(err.Error(), slog.String("method", info.FullMethod), slog.Any("req", reqStr))
		}

		return res, err
	}
}

// maskSensitiveFields replaces sensitive field values with "****"
func maskSensitiveFields(reqStr string, fields []string) string {
	for _, field := range fields {
		if strings.Contains(reqStr, field) {
			reqStr = strings.ReplaceAll(
				reqStr,
				fmt.Sprintf("%s:\"%s\"", field, extractFieldValue(reqStr, field)),
				field+":\"****\"",
			)
		}
	}

	return reqStr
}

// Helper function to extract the field value from the request string
func extractFieldValue(reqStr string, field string) string {
	start := strings.Index(reqStr, field+":\"") + len(field) + 2
	end := strings.Index(reqStr[start:], "\"") + start
	if start > -1 && end > -1 {
		return reqStr[start:end]
	}

	return ""
}
