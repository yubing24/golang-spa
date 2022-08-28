package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yubing24/golang-spa/api"
)

const (
	ENV_SEVER_PORT   = "SERVER_PORT"
	ENV_SPA_BASE_DIR = "SPA_BASE_DIR"
	ENV_SPA_ENTRY    = "SPA_ENTRY"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func initEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("💀 must specify environment to start server")
	}
}

func main() {
	initEnv()

	apiConf := api.Config{}
	http.Handle("/api", api.NewApiServer(apiConf))

	spa := spaHandler{staticPath: os.Getenv(ENV_SPA_BASE_DIR), indexPath: os.Getenv(ENV_SPA_ENTRY)}
	spaRouter := mux.NewRouter()
	spaRouter.Path("/").Handler(spa)
	http.Handle("/", spa)

	log.Printf("🚀 service starting...")
	http.ListenAndServe(":8000", nil)
}
