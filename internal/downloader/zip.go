package downloader

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadAndZip(info *GitHubURLInfo, outputFileName string) error {
	zipFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("could not create zip file : %w", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	initialAPIURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s?ref=%s",
		info.Owner, info.Repo, info.Path, info.Branch)

	return processDirectory(initialAPIURL, zipWriter)

}

func processDirectory(apiURL string, zipWriter *zip.Writer) error {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "zora-cli")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("github API responded with status: %s", resp.Status)
	}

	var contents []GitHubContent
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		return err
	}

	for _, item := range contents {
		if item.Type == "file" {
			fmt.Printf("  -> Adding file: %s\n", item.Path)
			if err := addFileToZip(item, zipWriter); err != nil {
				return err
			}
		} else if item.Type == "dir" {
			fmt.Printf("  -> Entering directory: %s\n", item.Path)
			if err := processDirectory(item.URL, zipWriter); err != nil {
				return err
			}
		}
	}
	return nil

}

func addFileToZip(file GitHubContent, zipWriter *zip.Writer) error {
	resp, err := http.Get(file.DownloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	zipEntry, err := zipWriter.Create(file.Path)
	if err != nil {
		return err
	}
	_, err = io.Copy(zipEntry, resp.Body)
	return err
}
