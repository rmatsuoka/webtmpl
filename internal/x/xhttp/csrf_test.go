package xhttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCSRF(t *testing.T) {
	doTest := func(t *testing.T, method, origin string, handler http.Handler) {
		t.Helper()
		req := httptest.NewRequest(method, "/", nil)
		req.Header.Set(HeaderOrigin, origin)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)
	}

	okHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("CSRF Attacked!")
	})

	t.Run("no origins", func(t *testing.T) {
		tests := []struct {
			method  string
			origin  string
			handler http.Handler
		}{
			{"GET", "https://example.com", okHandler},
			{"HEAD", "https://example.com", okHandler},
			{"POST", "https://example.com", okHandler},
			{"DELETE", "https://example.com", okHandler},
			{"GET", "https://other.example.com", okHandler},
			{"HEAD", "https://other.example.com", okHandler},
			{"POST", "https://attacker.example.com", panicHandler},
			{"OPTIONS", "https://attacker.example.com", panicHandler},
		}

		for _, test := range tests {
			handler := CSRF(nil, test.handler)
			doTest(t, test.method, test.origin, handler)
		}
	})

	t.Run("with specific origins", func(t *testing.T) {
		tests := []struct {
			method  string
			origin  string
			handler http.Handler
		}{
			{"GET", "https://example.com", okHandler},
			{"POST", "https://example.com", okHandler},
			{"DELETE", "https://api.example.com", okHandler},
			{"GET", "https//other.example.com", okHandler},
			{"HEAD", "https//other.example.com", okHandler},
			{"POST", "https://attcker.example.com", panicHandler},
			{"OPTIONS", "https://attacker.example.com", panicHandler},
		}

		for _, test := range tests {
			handler := CSRF([]string{"https://example.com", "https://api.example.com"}, test.handler)
			doTest(t, test.method, test.origin, handler)
		}
	})
}
