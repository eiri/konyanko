package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eiri/konyanko/ent"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var (
	dbName string
	client *ent.Client

	rootCmd = &cobra.Command{
		Use:   "konyanko",
		Short: "A nyaa.si aggregator",
		Long:  `A comand line utility to manage an RSS feed aggregator for nyaa.si`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			uri := fmt.Sprintf("file:%s?mode=rwc&cache=shared&_fk=1", dbName)
			client, err = ent.Open("sqlite3", uri)
			return err
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			client.Close()
		},
	}
)

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	rootCmd.PersistentFlags().StringVar(&dbName, "db", "nyaa.db", "database name")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
