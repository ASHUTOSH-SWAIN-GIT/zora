package cmd

import (
	"fmt"
	"os"

	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zora",
	Short: "Zora is a CLI tool to download a specific folder from GitHub.",
	Long: `A fast and simple command-line tool written in Go to download a
specific folder from a public GitHub repository without cloning the entire project.`,
	Run: func(cmd *cobra.Command, args []string) {
		ascii := figlet4go.NewAsciiRender()

		renderStr, _ := ascii.Render("Zora")
		fmt.Print(renderStr)

		fmt.Println()

		fmt.Println("A fast and simple tool to download specific folders from GitHub.")
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("\nUsage:")
		fmt.Println("  zora download <github-folder-url> [flags]")
		fmt.Println("\nExample:")
		fmt.Println("  zora download https://github.com/spf13/cobra/tree/main/docs")
		fmt.Println("\nRun 'zora download --help' for more information.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Flags and configuration settings can be defined here.
}
