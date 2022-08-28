package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path, _ := filepath.Abs(r.URL.Path)
		path = filepath.Join("./app/dist/app", path)
		log.Printf("🚀 Serving files in directory:  %v", path)
		// redirect traffic to index.html (SPA)
		_, fileExists := os.Stat(path)
		if os.IsNotExist(fileExists) {
			http.ServeFile(w, r, filepath.Join("./app/dist/app", "index.html"))
		}
		http.FileServer(http.Dir("./app/dist/app")).ServeHTTP(w, r)
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🛸 Hello from Server"))
	})

	http.ListenAndServe(":8000", nil)
}
