package api

import (
	"net/http"
)

func helloWorld(conf Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🛸 Hello, world! -- From Server"))
	}
}
