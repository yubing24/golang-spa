package api

import (
	"net/http"
)

func HelloWorld(conf Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🛸 Hello from Server"))
	}
}
