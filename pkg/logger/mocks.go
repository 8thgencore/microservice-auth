package logger

import (
	"context"
	"log/slog"
)

// MockHandler is a mock implementation of slog.Handler
type MockHandler struct {
	Logs []string
}

// Enabled implements slog.Handler interface and always returns true
func (m *MockHandler) Enabled(context.Context, slog.Level) bool {
	return true
}

// Handle implements slog.Handler interface and stores the log message
func (m *MockHandler) Handle(_ context.Context, r slog.Record) error {
	m.Logs = append(m.Logs, r.Message)
	return nil
}

// WithAttrs implements slog.Handler interface
func (m *MockHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return m
}

// WithGroup implements slog.Handler interface
func (m *MockHandler) WithGroup(_ string) slog.Handler {
	return m
}

// NewMockLogger creates a new logger with the mock handler
func NewMockLogger() *slog.Logger {
	mockHandler := &MockHandler{}
	return slog.New(mockHandler)
}
