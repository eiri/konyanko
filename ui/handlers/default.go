package handlers

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/eiri/konyanko/ui/components"
)

//go:embed assets
var assets embed.FS

type AnimeService interface {
	Episodes(ctx context.Context) ([]components.Episode, error)
}

type DefaultHandler struct {
	AnimeService AnimeService
}

func New(as AnimeService) *DefaultHandler {
	return &DefaultHandler{
		AnimeService: as,
	}
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r)
}

func (h *DefaultHandler) Static() http.Handler {
	if _, ok := os.LookupEnv("KONYANKO_DEV"); ok {
		root := http.Dir("./ui/handlers/assets")
		return http.FileServer(root)
	}
	root, err := fs.Sub(assets, "assets")
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(root))
}

func (h *DefaultHandler) Get(w http.ResponseWriter, r *http.Request) {
	episodes, err := h.AnimeService.Episodes(r.Context())
	if err != nil {
		http.Error(w, "failed to get episodes", http.StatusInternalServerError)
		return
	}
	h.View(w, r, episodes)
}

func (h *DefaultHandler) View(w http.ResponseWriter, r *http.Request, episodes []components.Episode) {
	components.Cards(episodes).Render(r.Context(), w)
}
