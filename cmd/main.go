package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	
	"go-social-video-downloader/utils"
)

const (
	defaultDownloadDir = "downloads"
)

func main() {
	fmt.Println("Go Social Video Downloader")
	
	if len(os.Args) < 2 {
		log.Fatal("Please provide a video URL as an argument")
	}

	videoURL := os.Args[1]
	
	// Validate URL
	if !utils.IsValidURL(videoURL) {
		log.Fatal("Invalid video URL. Please provide a valid YouTube or Facebook URL")
	}

	// Get download path
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter download path (leave blank for default): ")
	downloadPath, _ := reader.ReadString('\n')
	downloadPath = strings.TrimSpace(downloadPath)
	if downloadPath == "" {
		downloadPath = defaultDownloadDir
	}

	// Create download directory if it doesn't exist
	if err := os.MkdirAll(downloadPath, 0755); err != nil {
		log.Fatalf("Failed to create download directory: %v", err)
	}

	// Get quality preference
	fmt.Print("Enter video quality (high, medium, low): ")
	quality, _ := reader.ReadString('\n')
	quality = strings.TrimSpace(strings.ToLower(quality))

	// Initialize video downloader
	downloader, err := utils.NewVideoDownloader(videoURL, downloadPath, quality)
	if err != nil {
		log.Fatalf("Error initializing downloader: %v", err)
	}

	// Start download process
	fmt.Printf("Downloading video to: %s\n", downloadPath)
	if err := downloader.Download(); err != nil {
		log.Fatalf("Download failed: %v", err)
	}

	fmt.Println("Video download completed successfully!")
}
