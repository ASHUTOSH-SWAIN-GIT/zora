package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/ASHUTOSH-SWAIN-GIT/zora/internal/downloader"

	"github.com/spf13/cobra"
)

var outputFileName string

var downloadCmd = &cobra.Command{
	Use:   "download [github-folder-url]",
	Short: "Downloads the folder from the provided GitHub URL.",
	Long: `Takes a full GitHub URL to a folder, downloads its contents recursively,
and packages them into a zip file.

Example:
zora download https://github.com/spf13/cobra/tree/main/docs`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		githubURL := args[0]

		// Display header
		fmt.Println("Zora - GitHub Folder Downloader")
		fmt.Println("=================================")
		fmt.Println()

		// Parse URL with better feedback
		fmt.Print("Parsing GitHub URL... ")
		urlInfo, err := downloader.ParseGithubURL(githubURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError: Invalid URL. %v\n", err)
			os.Exit(1)
		}
		fmt.Println("OK")

		// Display parsed information in a nice format
		fmt.Println("Repository Information:")
		fmt.Printf("  Owner: %s\n", urlInfo.Owner)
		fmt.Printf("  Repository: %s\n", urlInfo.Repo)
		fmt.Printf("  Branch: %s\n", urlInfo.Branch)
		fmt.Printf("  Path: %s\n", urlInfo.Path)
		fmt.Println()

		// Download with progress indication
		fmt.Printf("Downloading files to '%s'...\n", outputFileName)
		fmt.Print("  ")

		// Create a simple progress indicator
		progressChars := []string{"|", "/", "-", "\\"}
		progressIndex := 0

		// Start progress indicator in a goroutine
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					fmt.Printf("\r  %s Downloading...", progressChars[progressIndex])
					progressIndex = (progressIndex + 1) % len(progressChars)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}()

		// Perform the actual download
		err = downloader.DownloadAndZip(urlInfo, outputFileName)
		done <- true

		if err != nil {
			fmt.Printf("\r  Download failed: %v\n", err)
			os.Exit(1)
		}

		// Success message with file info
		fmt.Printf("\r  Download completed successfully!\n")
		fmt.Println()

		// Get file size for display
		if fileInfo, err := os.Stat(outputFileName); err == nil {
			size := fileInfo.Size()
			var sizeStr string
			if size < 1024 {
				sizeStr = fmt.Sprintf("%d B", size)
			} else if size < 1024*1024 {
				sizeStr = fmt.Sprintf("%.1f KB", float64(size)/1024)
			} else {
				sizeStr = fmt.Sprintf("%.1f MB", float64(size)/(1024*1024))
			}
			fmt.Printf("Output file: %s (%s)\n", outputFileName, sizeStr)
		} else {
			fmt.Printf("Output file: %s\n", outputFileName)
		}

		fmt.Println("\nAll done! Your GitHub folder has been downloaded successfully.")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&outputFileName, "output", "o", "download.zip", "Name of the output zip file")
}
