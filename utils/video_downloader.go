package utils

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// VideoDownloader handles video download operations
type VideoDownloader struct {
	Platform     string
	URL         string
	DownloadPath string
	Quality      string
}

// IsValidURL checks if the URL is from a supported platform
func IsValidURL(videoURL string) bool {
	parsedURL, err := url.Parse(videoURL)
	if err != nil {
		return false
	}
	return getPlatformFromURL(parsedURL.Host) != ""
}

// NewVideoDownloader creates a new VideoDownloader instance
func NewVideoDownloader(videoURL, downloadPath, quality string) (*VideoDownloader, error) {
	parsedURL, err := url.Parse(videoURL)
	if err != nil {
		return nil, errors.New("invalid video URL")
	}

	platform := getPlatformFromURL(parsedURL.Host)
	if platform == "" {
		return nil, errors.New("unsupported video platform")
	}

	return &VideoDownloader{
		Platform:     platform,
		URL:          videoURL,
		DownloadPath: downloadPath,
		Quality:      quality,
	}, nil
}

// Download initiates the video download process
func (vd *VideoDownloader) Download() error {
	fmt.Printf("Downloading video from %s...\n", vd.Platform)

	switch vd.Platform {
	case "YouTube":
		return vd.downloadYouTubeVideo()
	case "Facebook":
		return vd.downloadFacebookVideo()
	default:
		return fmt.Errorf("platform %s not yet implemented", vd.Platform)
	}
}

// downloadYouTubeVideo handles YouTube video downloads using yt-dlp
func (vd *VideoDownloader) downloadYouTubeVideo() error {
	if _, err := exec.LookPath("yt-dlp"); err != nil {
		return errors.New("yt-dlp not found. Please install it first: https://github.com/yt-dlp/yt-dlp")
	}

	args := []string{
		"-o", filepath.Join(vd.DownloadPath, "%(title)s.%(ext)s"),
		"--restrict-filenames",
	}

	// Add quality selection
	switch vd.Quality {
	case "high":
		args = append(args, "-f", "best")
	case "medium":
		args = append(args, "-f", "best[height<=720]")
	case "low":
		args = append(args, "-f", "best[height<=480]")
	default:
		args = append(args, "-f", "best")
	}

	args = append(args, vd.URL)

	cmd := exec.Command("yt-dlp", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download video: %v", err)
	}

	return nil
}

// downloadFacebookVideo handles Facebook video downloads using yt-dlp
func (vd *VideoDownloader) downloadFacebookVideo() error {
	if _, err := exec.LookPath("yt-dlp"); err != nil {
		return errors.New("yt-dlp not found. Please install it first: https://github.com/yt-dlp/yt-dlp")
	}

	args := []string{
		"-o", filepath.Join(vd.DownloadPath, "%(title)s.%(ext)s"),
		"--restrict-filenames",
	}

	// Add quality selection
	switch vd.Quality {
	case "high":
		args = append(args, "-f", "best")
	case "medium":
		args = append(args, "-f", "best[height<=720]")
	case "low":
		args = append(args, "-f", "best[height<=480]")
	default:
		args = append(args, "-f", "best")
	}

	args = append(args, vd.URL)

	cmd := exec.Command("yt-dlp", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download video: %v", err)
	}

	return nil
}

// getPlatformFromURL determines the platform from the URL host
func getPlatformFromURL(host string) string {
	switch {
	case strings.Contains(host, "youtube.com"), strings.Contains(host, "youtu.be"):
		return "YouTube"
	case strings.Contains(host, "facebook.com"), strings.Contains(host, "fb.watch"):
		return "Facebook"
	default:
		return ""
	}
}
