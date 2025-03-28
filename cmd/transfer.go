package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfers a playlist from Spotify to another streaming service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Transferring songs...")
	},
}

func init() {
	rootCmd.AddCommand(transferCmd) // Attach `transfer` as a subcommand of `rootCmd`
}
