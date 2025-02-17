package api

import (
	"net/http"
	"sync"

	"github.com/rmatsuoka/webtmpl/internal/x/xhttp"
)

var (
	c  int
	mu sync.RWMutex
)

type countBody struct {
	Count int `json:"count"`
}

func count(w http.ResponseWriter, _ *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	xhttp.WriteJSON(w, 200, countBody{Count: c})
}

func countup(w http.ResponseWriter, _ *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	c++
	xhttp.WriteJSON(w, 200, countBody{Count: c})
}
