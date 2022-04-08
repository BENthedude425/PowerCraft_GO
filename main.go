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

	// Add handlers for paths

	mux.Handle("/", http.FileServer(http.FS(ui)))

	// Start the webserver
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
