package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var config Config

func init() {
	config = CreateConfig()
	fmt.Printf("CrudHost: %v\n", config.CRUDHost)
	fmt.Printf("CrudPort: %v\n", config.CRUDPort)
	fmt.Printf("Listening and Serving on Port: %v\n", config.ListenServePort)
}

func CreateConfig() Config {
	conf := Config{
		CRUDHost:        os.Getenv("CRUD_Host"),
		CRUDPort:        os.Getenv("CRUD_Port"),
		ListenServePort: os.Getenv("LISTEN_SERVE_PORT"),
	}
	return conf
}

func main() {
	server := Server{
		router: mux.NewRouter(),
	}
	//Set up routes for server
	server.routes()
	handler := removeTrailingSlash(server.router)
	fmt.Printf("starting server on port " + config.ListenServePort + "...\n")
	log.Fatal(http.ListenAndServe(":"+config.ListenServePort, handler))
}
func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
