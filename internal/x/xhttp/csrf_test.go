package xhttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCSRF(t *testing.T) {
	okHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	doTest := func(t *testing.T, wantStatusCode int, method, origin string, handler http.Handler) {
		t.Helper()
		req := httptest.NewRequest(method, "/", nil)
		req.Header.Set(HeaderOrigin, origin)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)
		got := w.Result().StatusCode
		if got != wantStatusCode {
			t.Errorf("Request(method=%s [Origin: %s]) returns Response.StatusCode = %d, want %d",
				method, origin, got, wantStatusCode)
		}
	}

	t.Run("no origins", func(t *testing.T) {
		handler := CSRF(nil, okHandler)
		tests := []struct {
			method         string
			origin         string
			wantStatusCode int
		}{
			{"GET", "https://example.com", 200},
			{"HEAD", "https://example.com", 200},
			{"POST", "https://example.com", 200},
			{"DELETE", "https://example.com", 200},
			{"GET", "https://other.example.com", 200},
			{"HEAD", "https://other.example.com", 200},
			{"POST", "https://attacker.example.com", 400},
			{"OPTIONS", "https://attacker.example.com", 400},
		}

		for _, test := range tests {
			doTest(t, test.wantStatusCode, test.method, test.origin, handler)
		}
	})

	t.Run("with specific origins", func(t *testing.T) {
		handler := CSRF([]string{"https://example.com", "https://api.example.com"}, okHandler)
		tests := []struct {
			method         string
			origin         string
			wantStatusCode int
		}{
			{"GET", "https://example.com", 200},
			{"POST", "https://example.com", 200},
			{"DELETE", "https://api.example.com", 200},
			{"GET", "https//other.example.com", 200},
			{"HEAD", "https//other.example.com", 200},
			{"POST", "https://attcker.example.com", 400},
			{"OPTIONS", "https://attacker.example.com", 400},
		}

		for _, test := range tests {
			doTest(t, test.wantStatusCode, test.method, test.origin, handler)
		}
	})
}
