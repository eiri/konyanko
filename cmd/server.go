package main

import (
	"fmt"
	"net/http"
	"os"

	gh "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/spf13/cobra"

	"github.com/eiri/konyanko"
	"github.com/eiri/konyanko/ui"
)

var (
	port int

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run UI and API server",
		RunE:  serverRunner,
	}
)

func init() {
	importCmd.Flags().IntVarP(&port, "port", "p", 2315, "A port to run the server on")

	rootCmd.AddCommand(serverCmd)
}

func serverRunner(cmd *cobra.Command, args []string) error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5, "application/json"))

	if _, ok := os.LookupEnv("KONYANKO_DEV"); ok {
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowCredentials: false,
			MaxAge:           300,
		}))
	}

	srv := gh.NewDefaultServer(konyanko.NewSchema(client))
	router.Handle("/playground",
		playground.Handler("Item", "/graphql"),
	)
	router.Handle("/graphql", srv)
	router.Handle("/", ui.NewStaticHandler())

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		return err
	}
	return nil
}
