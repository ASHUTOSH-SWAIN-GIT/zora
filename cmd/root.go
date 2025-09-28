package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "zora",
	Version: "1.0.0",
	Short:   "A CLI tool to download a specific folder from a GitHub repository.",
	Long: `zora is a fast and efficient command-line tool
that allows you to download the contents of a specific folder from a public
GitHub repository and save it as a single .zip file.

This is perfect for when you only need a subdirectory from a large project
without having to clone the entire repository.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
