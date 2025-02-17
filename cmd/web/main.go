package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/rmatsuoka/webtmpl/internal/api"
	"github.com/rmatsuoka/webtmpl/internal/env"
	"github.com/rmatsuoka/webtmpl/internal/x/xhttp"
	"github.com/rmatsuoka/webtmpl/internal/x/xslog"
)

func main() {
	logger := slog.New(&xslog.Handler{
		Handler: slog.NewTextHandler(os.Stderr, nil),
	})
	slog.SetDefault(logger)

	for pat, h := range api.Handlers() {
		http.Handle(pat, h)
	}

	slog.Info("start to listen", "addr", env.APP_LISTEN_ADDR)
	err := http.ListenAndServe(env.APP_LISTEN_ADDR, xhttp.LogHandler(http.DefaultServeMux))
	slog.Error("listen and server", "error", err)
}
