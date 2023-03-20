package handlers

import (
	"net/http"
)

func WelcomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	}
}
