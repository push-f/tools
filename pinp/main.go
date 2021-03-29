package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
)

//go:embed static
var static embed.FS

func main() {
	dir, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(dir)))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := net.JoinHostPort("", port)
	fmt.Printf("Run on http://localhost%s ...\n", addr)
	http.ListenAndServe(addr, nil)
}
