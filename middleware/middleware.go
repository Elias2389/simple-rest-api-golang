package middleware

import (
	"log"
	"net/http"
)

func Log(f func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request %q, Method: %q", r.URL.Path, r.Method)
		f(w, r)
	}
}
