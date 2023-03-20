package handlers

import (
	"calgo-server/src/db"
	"net/http"

	"gorm.io/gorm"
)

var datasource *gorm.DB = db.GetInstance()

func WelcomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	}
}
