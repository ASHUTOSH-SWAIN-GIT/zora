package downloader

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
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

	// Collect all files first
	var allFiles []GitHubContent
	err = collectAllFiles(initialAPIURL, &allFiles)
	if err != nil {
		return err
	}

	// Download files concurrently but write to zip sequentially
	return downloadAndZipFiles(allFiles, zipWriter)
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

func collectAllFiles(apiURL string, allFiles *[]GitHubContent) error {
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
			*allFiles = append(*allFiles, item)
		} else if item.Type == "dir" {
			// Recursively collect files from subdirectories
			if err := collectAllFiles(item.URL, allFiles); err != nil {
				return err
			}
		}
	}
	return nil
}

func downloadAndZipFiles(files []GitHubContent, zipWriter *zip.Writer) error {
	// Download files concurrently
	type fileResult struct {
		content []byte
		path    string
		err     error
	}

	results := make(chan fileResult, len(files))
	var wg sync.WaitGroup

	// Limit concurrent downloads
	semaphore := make(chan struct{}, 10)

	for _, file := range files {
		wg.Add(1)
		go func(f GitHubContent) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			content, err := downloadFileContent(f)
			results <- fileResult{content: content, path: f.Path, err: err}
		}(file)
	}

	// Close results channel when all downloads complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Write files to zip sequentially as they complete
	for result := range results {
		if result.err != nil {
			return fmt.Errorf("failed to download %s: %w", result.path, result.err)
		}

		zipEntry, err := zipWriter.Create(result.path)
		if err != nil {
			return fmt.Errorf("failed to create zip entry for %s: %w", result.path, err)
		}

		_, err = zipEntry.Write(result.content)
		if err != nil {
			return fmt.Errorf("failed to write %s to zip: %w", result.path, err)
		}
	}

	return nil
}

func downloadFileContent(file GitHubContent) ([]byte, error) {
	resp, err := http.Get(file.DownloadUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
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
