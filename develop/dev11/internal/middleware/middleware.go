package middleware

import (
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Req: ", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
