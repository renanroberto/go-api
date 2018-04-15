package router

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		duration := time.Since(start)

		log.Printf("%s \t %s \t %s\n", r.Method, r.RequestURI, duration)
	})
}
