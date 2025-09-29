# Zora

A fast and simple command-line tool written in Go to download specific folders from public GitHub repositories without cloning the entire project.

## Features

- **Fast Downloads**: Concurrent file downloads for maximum speed
- **Selective Download**: Download only the folder you need, not the entire repository
- **ZIP Packaging**: Automatically packages downloaded files into a zip archive
- **Beautiful Interface**: Clean ASCII art and progress indicators
- **Cross-Platform**: Works on Linux, macOS, and Windows

## Installation

### Quick Install (Recommended)

Install directly with Go:

```bash
go install github.com/ASHUTOSH-SWAIN-GIT/zora@latest
```

If `zora` is not found after installation, add Go's bin directory to your PATH:

```bash
# For Linux/macOS
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc  # or ~/.bashrc
source ~/.zshrc

# For Windows (PowerShell)
$env:PATH += ";$env:GOPATH\bin"
```

### Alternative: Download Binary

1. Download the latest release from the [Releases page](https://github.com/ASHUTOSH-SWAIN-GIT/zora/releases)
2. Extract the binary to your system PATH
3. Make it executable:
   ```bash
   chmod +x zora
   ```

## Usage

### Basic Usage

```bash
zora download <github-folder-url>
```

### Examples

Download a specific folder from a GitHub repository:

```bash
# Download the docs folder from cobra
zora download https://github.com/spf13/cobra/tree/main/docs

# Download components folder
zora download https://github.com/facebook/react/tree/main/packages/react-dom/src

# Download with custom output name
zora download https://github.com/microsoft/vscode/tree/main/extensions -o vscode-extensions.zip
```

### Command Options

```bash
zora download [github-folder-url] [flags]

Flags:
  -o, --output string   Name of the output zip file (default "download.zip")
  -h, --help           Help for download command
```

## Use Cases

- **Download Documentation**: Get just the docs folder from a project
- **Extract Components**: Download specific UI components or modules
- **Sample Code**: Get example code without the entire repository
- **Configuration Files**: Download config templates or examples
- **Assets**: Download images, fonts, or other static assets

## How It Works

1. **URL Parsing**: Extracts repository owner, name, branch, and folder path from GitHub URL
2. **API Discovery**: Uses GitHub API to discover all files in the specified folder
3. **Concurrent Download**: Downloads files concurrently (up to 10 simultaneous downloads)
4. **ZIP Creation**: Packages all downloaded files into a single zip archive
5. **Progress Feedback**: Shows real-time download progress

## Development

### Prerequisites

- Go 1.19 or higher
- Git

### Building

```bash
# Clone the repository
git clone https://github.com/ASHUTOSH-SWAIN-GIT/zora.git
cd zora

# Install dependencies
go mod tidy

# Build the binary
go build -o zora

# Run tests
go test ./...
```

### Project Structure

```
zora/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command with ASCII art
│   └── download.go        # Download command implementation
├── internal/
│   └── downloader/        # Core download logic
│       ├── github.go      # GitHub API integration
│       ├── parser.go      # URL parsing
│       └── zip.go         # ZIP file creation
├── main.go                # Application entry point
├── go.mod                 # Go module definition
└── README.md              # This file
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI framework
- Uses [figlet4go](https://github.com/mbndr/figlet4go) for ASCII art
- Inspired by the need for selective GitHub folder downloads

## Troubleshooting

### Common Issues

**"github API responded with status: 404 Not Found"**
- Make sure the repository is public
- Check that the folder path exists
- Verify the branch name is correct

**"Invalid URL"**
- Ensure you're using the full GitHub URL to a folder (not a file)
- URL should be in format: `https://github.com/owner/repo/tree/branch/path`

**Permission denied when installing**
- Use `sudo` for system-wide installation
- Or install to user directory with `~/.local/bin`

### Getting Help

- Check the [Issues](https://github.com/ASHUTOSH-SWAIN-GIT/zora/issues) page
- Create a new issue if you don't find your problem
- Include your operating system and the full error message

---

Made with love by [ASHUTOSH-SWAIN-GIT](https://github.com/ASHUTOSH-SWAIN-GIT)
