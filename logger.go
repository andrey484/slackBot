package main

import (
	"net/http"
	"time"
	"log"
)

func MyLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		start := time.Now()
		inner.ServeHTTP(response, request)

		log.Printf("%s\t%s\t%s\t%s",
			request.Method,
			request.RequestURI,
			name,
			time.Since(start))
	})
}
