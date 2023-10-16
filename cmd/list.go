package main

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print database records",
}

func init() {
	rootCmd.AddCommand(listCmd)
}
