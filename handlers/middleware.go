package handlers

import (
	"net/http"
)

var (
	apiKeyHdr = "X-Api-Key"
	apiKey    = "devkey"
)

func (h *Handler) MiddlewareAuthz(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// authorize request
		k := r.Header.Get(apiKeyHdr)
		if ok := verifyAPIKey(k); !ok {
			http.Error(w, "error not authorized", http.StatusForbidden)
			return
		}

		// next handler
		next.ServeHTTP(w, r)
	})
}

func verifyAPIKey(k string) bool {
	return k == apiKey
}

func isValidAddr(addr string) error {
	return nil
}

func areValidAddrs(addrs string) error {
	return nil
}

func areValidTimes(min, max string) error {
	return nil
}
