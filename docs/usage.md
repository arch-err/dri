# Usage

## Starting dri

```bash
dri
```

This launches the interactive TUI. You'll see a list of all repositories synced with your Drone CI instance.

## Navigation Flow

```
Repositories → Builds → Log Viewer
```

### Repository List

- Scroll through repositories with arrow keys or `j`/`k`
- Press `/` to fuzzy search by repository name
- Press `enter` to view builds for the selected repository

### Build List

- Browse builds with status indicators, event type, branch, and author
- Press `/` to filter builds
- Press `enter` to view build logs
- Press `esc` to go back to repositories

### Log Viewer

- Logs are displayed in a tabbed interface with one tab per build step
- Use `tab` and `shift+tab` to switch between steps
- Scroll with arrow keys, `pgup`/`pgdn`, or `home`/`end`
- ANSI colors from build output are preserved
- Press `esc` to go back to the build list

## Version

```bash
dri --version
```
