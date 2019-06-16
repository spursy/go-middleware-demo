package main 

import (
	"github.com/goji/httpauth"
	"net/http"
)

func main() {
	finalHandler := http.HandlerFunc(final)
	authHandler := httpauth.SimpleBasicAuth("username", "password")

	http.Handle("/", authHandler(finalHandler))
	http.ListenAndServe(":3000", nil)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
