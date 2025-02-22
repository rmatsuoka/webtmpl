package env

import (
	"cmp"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/rmatsuoka/webtmpl/internal/x/must"
	"github.com/rmatsuoka/webtmpl/internal/x/xslog"
)

var getenv = os.Getenv

var (
	APP_LISTEN_ADDR      = envString("APP_LISTEN_ADDR", ":8080")
	APP_LOG_LEVEL        = envFunc("APP_LOG_LEVEL", slog.LevelInfo, xslog.ParseLevel)
	APP_SHUTDOWN_TIMEOUT = envFunc("APP_SHUTDOWN_TIMEOUT", time.Second*10, time.ParseDuration)
	APP_CSRF_ORIGINS     = envStringSlice("APP_CSRF_ORIGINS", nil)
)

func envString(key string, fallback string) string {
	return cmp.Or(getenv(key), fallback)
}

func envStringSlice(key string, fallback []string) []string {
	v := getenv(key)
	if v == "" {
		return fallback
	}
	return strings.Split(v, ",")
}

func envFunc[T any](key string, fallback T, parse func(string) (T, error)) T {
	v := getenv(key)
	if v == "" {
		return fallback
	}
	return must.Do(parse(v))
}
