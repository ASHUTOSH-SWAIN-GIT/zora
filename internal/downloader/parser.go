package downloader

import (
	"fmt"
	"net/url"
	"strings"
)

type GitHubURLInfo struct {
	Owner  string
	Repo   string
	Branch string
	Path   string
}

func ParseGithubURL(rawURL string) (*GitHubURLInfo, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid url format : %w", err)
	}
	if parsedURL.Host != "github.com" {
		return nil, fmt.Errorf("url must be a github.com link ")
	}

	cleanPath := strings.Trim(parsedURL.Path, "/")
	parts := strings.Split(cleanPath, "/")

	if len(parts) < 4 || parts[2] != "tree" {
		return nil, fmt.Errorf("invalid or unsupported GitHub folder URL. Expecting the format '/owner/repo/tree/branch/path'")
	}

	info := &GitHubURLInfo{
		Owner:  parts[0],
		Repo:   parts[1],
		Branch: parts[3],
		Path:   strings.Join(parts[4:], "/"),
	}
	return info, nil
}
