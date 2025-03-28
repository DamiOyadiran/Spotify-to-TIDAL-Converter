package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "spotify-transfer"
	Short: "CLI tool to transfer songs from Spotify to another streaming service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Spotify Transfer CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}