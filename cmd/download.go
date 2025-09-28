package cmd

import (
	"fmt"
	"os"
	"zora/internal/downloader"

	"github.com/spf13/cobra"
)

var outputFileName string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download [github-folder-url]",
	Short: "Downloads the folder from the provided GitHub URL.",
	Long: `Takes a full GitHub URL to a folder, downloads its contents recursively,
and packages them into a zip file.

Example:
zora download https://github.com/spf13/cobra/tree/main/docs`,
	// Enforce that exactly one argument (the URL) is provided.
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		githubURL := args[0]

		// 1. Parse the provided URL to extract its components.
		fmt.Println("-> Parsing GitHub URL...")
		urlInfo, err := downloader.ParseGitHubURL(githubURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid URL. %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✓ URL Parsed: [Owner: %s, Repo: %s, Branch: %s, Path: %s]\n",
			urlInfo.Owner, urlInfo.Repo, urlInfo.Branch, u.rlInfo.Path)

		// 2. Start the download and zip process.
		fmt.Printf("-> Starting download to '%s'...\n", outputFileName)
		err = downloader.DownloadAndZip(urlInfo, outputFileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Download failed. %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\n✅ Success! Folder downloaded and saved to '%s'.\n", outputFileName)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Add the --output (-o) flag to allow users to specify the output file name.
	downloadCmd.Flags().StringVarP(&outputFileName, "output", "o", "download.zip", "Name of the output zip file")
}
