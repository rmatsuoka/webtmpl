package xhttp

import (
	"cmp"
	"log/slog"
	"net/http"

	"github.com/rmatsuoka/webtmpl/internal/x/xslog"
)

func LogHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := xslog.ContextWith(req.Context(), slog.Group("request",
			"method", req.Method,
			"requestURI", req.URL.RequestURI(),
			"remoteAddr", req.RemoteAddr,
			"userAgent", req.UserAgent(),
		))
		lw := &logResponseWriter{ResponseWriter: w}
		h.ServeHTTP(lw, req.WithContext(ctx))

		slog.InfoContext(ctx, "request",
			"statusCode", lw.StatusCode(),
		)
	})
}

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *logResponseWriter) StatusCode() int {
	return cmp.Or(w.statusCode, 200)
}

func (w *logResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *logResponseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

var _ http.ResponseWriter = (*logResponseWriter)(nil)
