# Go Social Video Downloader

A command-line tool for downloading videos from YouTube and Facebook.

## Features
- Download videos from YouTube and Facebook
- Select video quality (high, medium, low)
- Specify custom download directory
- Portable static binary

## Installation

### From Source
1. Ensure you have Go installed (version 1.16 or higher)
2. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/go-social-video-downloader.git
   cd go-social-video-downloader
   ```
3. Build the application:
   ```bash
   make static-build
   ```
4. The binary will be created in the `bin/` directory

### Using Pre-built Binary
1. Download the latest release from the [Releases page](#)
2. Make the binary executable:
   ```bash
   chmod +x go-social-video-downloader
   ```
3. Move it to your PATH:
   ```bash
   sudo mv go-social-video-downloader /usr/local/bin/
   ```

## Usage

Basic usage:
```bash
go-social-video-downloader [video_url]
```

Example:
```bash
go-social-video-downloader https://www.youtube.com/watch?v=dQw4w9WgXcQ
```

### Options
- When prompted for download path, press Enter to use default (`downloads/`)
- When prompted for quality, enter `high`, `medium`, or `low`

## Requirements
- yt-dlp installed and available in PATH

## License
MIT License

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Disclaimer
This tool is for educational purposes only. Please respect copyright laws and only download videos you have the rights to.
