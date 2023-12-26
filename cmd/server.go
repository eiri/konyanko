package main

import (
	"fmt"
	"net/http"

	gh "github.com/99designs/gqlgen/graphql/handler"
	"github.com/spf13/cobra"

	"github.com/eiri/konyanko"
	"github.com/eiri/konyanko/ui/handlers"
	"github.com/eiri/konyanko/ui/services"
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
	srv := gh.NewDefaultServer(konyanko.NewSchema(client))
	// http.Handle("/playground",
	// 	playground.Handler("Item", "/graphql"),
	// )
	http.Handle("/graphql", srv)

	as := services.NewAnime(client)
	h := handlers.New(as)
	http.Handle("/", h.Static())
	http.Handle("/list", h)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		return err
	}
	return nil
}
