package env

import (
	"cmp"
	"os"
)

var (
	APP_LISTEN_ADDR = envString("APP_LISTEN_ADDR", ":8080")
)

func envString(key string, fallback string) string {
	return cmp.Or(os.Getenv(key), fallback)
}
