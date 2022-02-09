package sqlxplus

import (
	"context"
	"log"
	"strings"
)

// Logger logger
type Logger interface {
	// Printf 打印
	Printf(ctx context.Context, query string, args ...interface{})
	// Println 打印
	Println(ctx context.Context, args ...interface{})
	// Errorf 错误
	Errorf(ctx context.Context, format string, args ...interface{})
	// Errorln 错误
	Errorln(ctx context.Context, args ...interface{})
}

// StdLogger ...
type StdLogger struct {
}

// Printf 打印
func (StdLogger) Printf(ctx context.Context, query string, args ...interface{}) {
	query = strings.ReplaceAll(query, "?", "%v")
	log.Printf("[sqlx] "+query, args...)
}

// Println 打印
func (StdLogger) Println(ctx context.Context, args ...interface{}) {
	nArgs := []interface{}{
		"[sqlx]",
	}
	nArgs = append(nArgs, args...)
	log.Println(nArgs...)
}

// Errorf 错误
func (StdLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Printf("[sqlx-error] "+format, args...)
}

// Errorln 错误
func (StdLogger) Errorln(ctx context.Context, args ...interface{}) {
	nArgs := []interface{}{
		"[sqlx-error]",
	}
	nArgs = append(nArgs, args...)
	log.Println(nArgs...)
}
