package ui

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed dist
var dist embed.FS

func NewStaticHandler() http.Handler {
	if _, ok := os.LookupEnv("KONYANKO_DEV"); ok {
		root := http.Dir("./ui/dist")
		return http.FileServer(root)
	}
	root, err := fs.Sub(dist, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(root))
}
