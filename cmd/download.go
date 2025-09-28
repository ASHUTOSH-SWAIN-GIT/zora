package cmd

import "github.com/spf13/cobra"

var outputFile string

var downloadCmd = &cobra.Command{
	Use:   "download [github-folder-url]",
	Short: "Downloads the folders from the provided github url",
	Long: `Takes a full GitHub URL to a folder, downloads its contents recursively,
and packages them into a zip file.

Example:
zora download https://github.com/spf13/cobra/tree/main/docs`,
}
Args: cobra.ExactArgs(1),
