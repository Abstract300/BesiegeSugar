package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := new(http.ServeMux)
	server := makeServer()
	server.Handler = mux
	defer server.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", "Home page")
	})

	mux.HandleFunc("/", siteHandler)

	server.ListenAndServe()
}

func makeServer() *http.Server {
	return &http.Server{
		Addr: "localhost:8080",
	}
}

func siteHandler(w http.ResponseWriter, r *http.Request) {
	post := r.URL.Path[len("/blog/"):]
	io.WriteString(w, post)
}
