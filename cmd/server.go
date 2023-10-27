package main

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/cobra"

	"github.com/eiri/konyanko"
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
	srv := handler.NewDefaultServer(konyanko.NewSchema(client))
	http.Handle("/playground",
		playground.Handler("Item", "/query"),
	)
	http.Handle("/query", srv)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		return err
	}
	return nil
}
