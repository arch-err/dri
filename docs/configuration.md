# Configuration

dri is configured through environment variables.

## Required Variables

| Variable | Description |
|----------|-------------|
| `DRONE_SERVER` | URL of your Drone CI server (e.g., `https://drone.example.com`) |
| `DRONE_TOKEN` | Your Drone API token |

## Getting Your Token

1. Log into your Drone CI instance
2. Navigate to your account settings
3. Copy the API token shown on the page

## Shell Configuration

Add to your shell profile (`~/.bashrc`, `~/.zshrc`, etc.):

```bash
export DRONE_SERVER=https://drone.example.com
export DRONE_TOKEN=your-api-token
```
