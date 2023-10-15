package main

import (
	"fmt"
	"net/http"

	"github.com/eiri/konyanko/ent/ogent"
	"github.com/spf13/cobra"
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
	srv, err := ogent.NewServer(ogent.NewOgentHandler(client))
	if err != nil {
		return err
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), srv); err != nil {
		return err
	}
	return nil
}
