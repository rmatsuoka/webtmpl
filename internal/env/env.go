package env

import (
	"cmp"
	"encoding"
	"log/slog"
	"os"
)

var (
	APP_LISTEN_ADDR = envString("APP_LISTEN_ADDR", ":8080")
	APP_LOG_LEVEL   = *envAs[*slog.Level]("APP_LOG_LEVEL", "INFO")
)

func envString(key string, fallback string) string {
	return cmp.Or(os.Getenv(key), fallback)
}

func envAs[T encoding.TextUnmarshaler](key string, fallback string) T {
	var t T
	t.UnmarshalText([]byte(cmp.Or(os.Getenv(key), fallback)))
	return t
}
