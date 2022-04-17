package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

//go:embed all:ui
var ui embed.FS

const configFilePath = "config.txt"

type ConfigPresets struct {
	dirofsoemthing string
}

func Initialise() {
	_, err := os.OpenFile(configFilePath, os.O_RDONLY, 0666)
	if err != nil {

		log.Output(1, "No config file detected.... Creating new with default config")

		configFile, err2 := os.Create(configFilePath)

		if err2 != nil {
			log.Fatal(err2)
		}
		_, err3 := configFile.Write([]byte(""))

		if err3 != nil {
			log.Fatal(err3)
		}

	}

}

func CreateUser(username string, password string) {

}

func LoginUser(username string, password string) {
	println(username, password)
}

func main() {
	Initialise()
	// Parse command-line flags
	flagVersion := flag.Bool("build", false, "print build information then exit")
	flag.Parse()

	// Handle build flag
	if *flagVersion {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("build info not found")
			os.Exit(1)
		}
		fmt.Printf("%v", info)
		os.Exit(0)
	}

	// Remove /ui/ prefix from ui embed FS
	ui, _ := fs.Sub(ui, "ui")

	// Create a HTTP multiplexer to handle requests to paths
	mux := http.NewServeMux()

	//
	// Add handlers for paths
	//

	// Handle requests to the API
	mux.HandleFunc("/api/", func(w http.ResponseWriter, req *http.Request) {
		// The "/api" pattern matches everything, so we handle specific
		// paths here that we're at the root here.
		println(req.Method, req.URL.Path) // Displays the method and URL for debugging.
		if req.Method == "GET" {
			if req.URL.Path == "/api/test" {
				fmt.Fprintf(w, "test successful")
			} else {
				http.NotFound(w, req)
				return
			}
		} else if req.Method == "POST" {
			err := req.ParseForm()
			if err != nil {
				log.Fatal(err)
			}

			if req.URL.Path == "/api/login" {
				username := req.PostForm.Get("username")
				password := req.PostForm.Get("password")
				LoginUser(username, password)
			} else if req.URL.Path == "/api/create" {
				username := req.PostForm.Get("username")
				password := req.PostForm.Get("password")
				CreateUser(username, password)
			}
		}
	})

	// Handle requests to the root by returning the UI
	mux.Handle("/", http.FileServer(http.FS(ui)))

	// Start the webserver
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		panic(err)
	}
}
