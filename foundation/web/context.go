package web

import (
	"context"
	"time"
)

type ctxKey int

const key ctxKey = 1

type Values struct {
	TraceId    string
	StatusCode int
	Now        time.Time
}

// Returns values from context
func GetValues(ctx context.Context) *Values {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return &Values{
			TraceId: "00000000-0000-0000-0000-000000000000",
			Now:     time.Now(),
		}
	}
	return v
}

// Returns traceId from context
func GetTraceId(ctx context.Context) string {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return "00000000-0000-0000-0000-000000000000"
	}
	return v.TraceId
}

// Returns time from context
func GetTime(ctx context.Context) time.Time {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return time.Now()
	}
	return v.Now
}

// Sets status in context
func SetStatus(ctx context.Context, statusCode int) {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return
	}
	v.StatusCode = statusCode
}
