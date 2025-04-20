# Sims 4 Mod Manager

A version-controlled mod management CLI for The Sims 4 with a cute themed UI!

<img src="/api/placeholder/400/100" alt="Sims 4 Mod Manager Logo" />

## Features

- 📁 List all mod folders in your Sims 4 mods directory
- 🔄 Track mod versions for easy update management
- 💾 Store old versions of mods for safe rollbacks
- 🎮 Cute Sims 4 themed UI with the iconic plumbob colors
- 🖥️ Both CLI commands and interactive TUI interface

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.21 or higher)
- The Sims 4 installed with a Mods folder

### Install from source

1. Clone the repository:
   ```bash
   git clone https://github.com/naptalie/sims4-mod-manager.git
   cd sims4-mod-manager
   ```

2. Build the application:
   ```bash
   go build -o sims4-mod-manager
   ```

3. Move the binary to your PATH (optional):
   ```bash
   # On Linux/macOS
   sudo mv sims4-mod-manager /usr/local/bin/

   # On Windows, move to a directory in your PATH
   ```

## Quick Start

```bash
# Launch the interactive UI
sims4-mod-manager ui

# List all mods in your Sims 4 mods folder
sims4-mod-manager list

# Backup all mods
sims4-mod-manager backup

# Restore a specific mod to a previous version
sims4-mod-manager restore "ModName" "20250419_120000"
```

## Project Structure

```
sims4-mod-manager/
├── cmd/                # Command handling
│   ├── backup.go       # Backup command logic
│   ├── config.go       # Config command logic
│   ├── list.go         # List command logic
│   ├── restore.go      # Restore command logic
│   ├── root.go         # Root command definition
│   ├── ui.go           # UI command launcher
│   └── versions.go     # Versions command logic
├── internal/           # Internal packages
│   ├── config/         # Configuration handling
│   │   └── config.go   # Config management
│   ├── models/         # Data models
│   │   └── mod.go      # Mod data structure
│   ├── ui/             # UI components
│   │   ├── styles/     # UI styling
│   │   │   └── styles.go # Theme colors and styles
│   │   └── tui/        # Terminal UI
│   │       ├── backup.go  # Backup UI component
│   │       ├── modlist.go # Mod list component
│   │       ├── restore.go # Restore UI component
│   │       └── tui.go     # Main TUI component
│   └── core/           # Core functionality
│       ├── scanner.go  # Mod scanning logic
│       ├── backup.go   # Backup operations
│       └── restore.go  # Restore operations
├── pkg/                # Public packages
│   └── fsutil/         # File system utilities
│       └── fsutil.go   # File operations
├── main.go             # Application entry point
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── Makefile            # Build automation
└── README.md           # This file
```

## Development

### Adding New Features

1. Fork the repository
2. Create a feature branch
3. Add your changes
4. Run `go build` to test
5. Submit a pull request

### Building for Different Platforms

```bash
# Build for Windows
GOOS=windows GOARCH=amd64 go build -o sims4-mod-manager.exe

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o sims4-mod-manager-mac

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o sims4-mod-manager-linux
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- The Sims 4 and its assets are property of Electronic Arts/Maxis
- Built with [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper), and [Bubble Tea](https://github.com/charmbracelet/bubbletea)