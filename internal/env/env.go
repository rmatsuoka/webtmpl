package env

import (
	"cmp"
	"encoding"
	"log/slog"
	"os"
	"time"

	"github.com/rmatsuoka/webtmpl/internal/x/must"
)

var getenv = os.Getenv

var (
	APP_LISTEN_ADDR      = envString("APP_LISTEN_ADDR", ":8080")
	APP_LOG_LEVEL        = *envText("APP_LOG_LEVEL", "INFO", new(slog.Level))
	APP_SHUTDOWN_TIMEOUT = envFunc("APP_SHUTDOWN_TIMEOUT", time.Second*10, time.ParseDuration)
)

func envString(key string, fallback string) string {
	return cmp.Or(getenv(key), fallback)
}

func envText[T encoding.TextUnmarshaler](key string, fallback string, t T) T {
	must.Nil(t.UnmarshalText([]byte(cmp.Or(getenv(key), fallback))))
	return t
}

func envFunc[T any](key string, fallback T, parse func(string) (T, error)) T {
	v := getenv(key)
	if v == "" {
		return fallback
	}
	return must.Do(parse(v))
}
