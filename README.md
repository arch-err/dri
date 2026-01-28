# dri

**dri** (drone-interactive) is a terminal user interface for [Drone CI](https://www.drone.io/).

Browse repositories, inspect builds, and read step logs — all from your terminal.

## Features

- 🚀 Browse all your Drone CI repositories with fuzzy search
- 📊 View build history with status indicators
- 📜 Read build logs with tabbed step navigation
- 🎨 Full ANSI color support in log output
- ⌨️ Keyboard-driven navigation

## Installation

### Binary Releases

Download from the [releases page](https://github.com/arch-err/dri/releases).

### From Source

```bash
go install github.com/arch-err/dri/cmd/dri@latest
```

## Configuration

Set your Drone server URL and API token:

```bash
export DRONE_SERVER=https://drone.example.com
export DRONE_TOKEN=your-api-token
```

## Usage

```bash
dri
```

### Navigation

| Key | Action |
|-----|--------|
| `enter` | Select / drill down |
| `esc` | Go back |
| `/` | Filter list |
| `tab` / `shift+tab` | Switch log tabs |
| `↑` / `↓` / `pgup` / `pgdn` | Scroll |
| `q` / `ctrl+c` | Quit |

## Documentation

Full documentation available at [https://arch-err.github.io/dri](https://arch-err.github.io/dri)

## License

MIT
