# Contributing

## Development Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/arch-err/dri.git
    cd dri
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. Build:

    ```bash
    go build ./cmd/dri
    ```

4. Run:

    ```bash
    export DRONE_SERVER=https://drone.example.com
    export DRONE_TOKEN=your-token
    ./dri
    ```

## Project Structure

```
cmd/dri/             Entrypoint
internal/
  client/            Drone SDK wrapper
  config/            Environment configuration
  tui/               Bubbletea TUI
    builds/          Build list screen
    logs/            Log viewer screen
    msg/             Shared message types
    repos/           Repository list screen
    styles/          Shared lipgloss styles
  version/           Version variable (set via ldflags)
```
