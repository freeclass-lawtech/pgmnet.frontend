package main

import (
    "embed"
    "fmt"
    "net/http"
    "log"
)

//go:embed dist/*
var staticFiles embed.FS

func logRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Received request: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func main() {
    // Serve dist files with logging and MIME type handling
	http.Handle("/dist/", logRequest(http.StripPrefix("/dist/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		data, err := staticFiles.ReadFile("dist/" + path)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		// Set MIME type based on file extension
		if len(path) > 3 && path[len(path)-3:] == ".js" {
			w.Header().Set("Content-Type", "application/javascript")
		} else if len(path) > 4 && path[len(path)-4:] == ".css" {
			w.Header().Set("Content-Type", "text/css")
		}

		w.Write(data)
	}))))

    // Serve the main page
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data, err := staticFiles.ReadFile("dist/index.html")
        if err != nil {
            http.Error(w, "Could not read embedded file", http.StatusInternalServerError)
            return
        }
        fmt.Fprintln(w, string(data))
    })

    // Start the server
    log.Println("Starting server at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
