package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "train",
	Short: "get train info",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		return
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(list)
	rootCmd.AddCommand(station)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Demo",
	Long:  `All software has versions. This is demo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: v1.0")
	},
}
