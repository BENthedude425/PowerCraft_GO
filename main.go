package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:ui
var ui embed.FS

func main() {
	// Remove /ui/ prefix from ui embed FS
	ui, _ := fs.Sub(ui, "ui")

	// Create a HTTP multiplexer to handle requests to paths
	mux := http.NewServeMux()

	//
	// Add handlers for paths
	//

	// Handle requests to the API
	mux.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		// The "/api" pattern matches everything, so we handle specific
		// paths here that we're at the root here.
		if req.URL.Path == "/api/test" {
			fmt.Fprintf(w, "test successful",)
		} else {
			http.NotFound(w, req)
			return
		}
	})

	// Handle requests to the root by returning the UI
	mux.Handle("/", http.FileServer(http.FS(ui)))

	// Start the webserver
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
