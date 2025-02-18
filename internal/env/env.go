package env

import (
	"cmp"
	"encoding"
	"log/slog"
	"os"

	"github.com/rmatsuoka/webtmpl/internal/x/must"
)

var getenv = os.Getenv

var (
	APP_LISTEN_ADDR = envString("APP_LISTEN_ADDR", ":8080")
	APP_LOG_LEVEL   = *envAs[*slog.Level]("APP_LOG_LEVEL", "INFO")
)

func envString(key string, fallback string) string {
	return cmp.Or(getenv(key), fallback)
}

func envAs[T encoding.TextUnmarshaler](key string, fallback string) T {
	var t T
	must.Nil(t.UnmarshalText([]byte(cmp.Or(getenv(key), fallback))))
	return t
}
