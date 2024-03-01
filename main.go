package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	PORT := envPortOr("3000")
	}
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	routes.Route()
	
	fmt.Printf("Starting server on port %s...\n", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
