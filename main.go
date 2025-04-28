package main

import "net/http"

func main() {
	// Serve static files from "./static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
