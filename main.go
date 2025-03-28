package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spotify-transfer",
	Short: "Transfer Spotify playlists between accounts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Spotify Transfer CLI")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
