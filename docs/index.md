# dri

**dri** (drone-interactive) is a terminal user interface for [Drone CI](https://www.drone.io/). Browse repositories, inspect builds, and read step logs — all from your terminal.

## Features

- Browse all your Drone CI repositories with fuzzy search
- View build history with status indicators
- Read build logs with tabbed step navigation
- Full ANSI color support in log output
- Keyboard-driven navigation

## Quick Start

```bash
export DRONE_SERVER=https://drone.example.com
export DRONE_TOKEN=your-api-token
dri
```

## Navigation

| Key | Action |
|-----|--------|
| `enter` | Select / drill down |
| `esc` | Go back |
| `/` | Filter list |
| `tab` / `shift+tab` | Switch log tabs |
| `q` / `ctrl+c` | Quit |
