package database

import (
	"context"
	"log"
	"time"
)

const startTimeKey = "psql-start"

type QueryLogger struct{}

func (h *QueryLogger) Before(ctx context.Context, query string, args ...any) (context.Context, error) {
	// log.Printf("[SQL BEFORE] %s | %v", query, args)
	// return ctx, nil
	log.Printf("[SQL START] %s | args=%v", query, args)
	start := time.Now()
	return context.WithValue(ctx, startTimeKey, start), nil
}

func (h *QueryLogger) After(ctx context.Context, query string, args ...any) (context.Context, error) {
	// log.Printf("[SQL AFTER] %s | %v", query, args)
	if start, ok := ctx.Value(startTimeKey).(time.Time); ok {
		log.Printf("[SQL OK] %s | duration=%s", query, time.Since(start))
	} else {
		log.Printf("[SQL OK] %s", query)
	}

	return ctx, nil
}

func (h *QueryLogger) OnError(ctx context.Context, err error, query string, args ...any) error {
	// log.Printf("[SQL ERROR] %s | %v | err=%v", query, args, err)
	if start, ok := ctx.Value(startTimeKey).(time.Time); ok {
		log.Printf("[SQL FAIL] %s | args=%v | duration=%s | err=%v", query, args, time.Since(start), err)
	} else {
		log.Printf("[SQL FAIL] %s | args=%v | err=%v", query, args, err)
	}
	return err
}
