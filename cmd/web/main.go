package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rmatsuoka/webtmpl"
	"github.com/rmatsuoka/webtmpl/internal/api"
	"github.com/rmatsuoka/webtmpl/internal/env"
	"github.com/rmatsuoka/webtmpl/internal/x/xhttp"
	"github.com/rmatsuoka/webtmpl/internal/x/xslog"
)

func main() {
	sigctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	logger := slog.New(&xslog.Handler{
		Handler: slog.NewTextHandler(os.Stderr, nil),
	})
	slog.SetDefault(logger)

	for pat, h := range api.Handlers() {
		http.Handle(pat, xhttp.LogHandler(h))
	}
	http.Handle("GET /statics/", http.FileServerFS(webtmpl.Content()))
	http.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFileFS(w, req, webtmpl.Content(), "index.html")
	})

	srv := &http.Server{
		Addr: env.APP_LISTEN_ADDR,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		<-sigctx.Done()
		ctx, stop := context.WithTimeout(context.Background(), time.Second*10)
		defer stop()

		err := srv.Shutdown(ctx)
		if err != nil {
			slog.Error("server shutdown", "error", err)
		}
		close(idleConnsClosed)
	}()

	slog.Info("start to listen", "addr", env.APP_LISTEN_ADDR)
	err := srv.ListenAndServe()
	slog.Error("listen and server", "error", err)
	<-idleConnsClosed
}
