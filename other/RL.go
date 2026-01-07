package main

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

func limitMiddleware(next http.Handler) http.Handler {
	// 10 запросов в секунду
	limiter := rate.NewLimiter(10, 20)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	limitedMux := limitMiddleware(mux)
	http.ListenAndServe(":8080", limitedMux)
}
