# Installation

## Binary Releases

Download the latest binary from the [releases page](https://github.com/arch-err/dri/releases).

=== "Linux (amd64)"

    ```bash
    curl -Lo dri https://github.com/arch-err/dri/releases/latest/download/dri_linux_amd64.tar.gz
    tar xzf dri_linux_amd64.tar.gz
    sudo mv dri /usr/local/bin/
    ```

=== "macOS (Apple Silicon)"

    ```bash
    curl -Lo dri https://github.com/arch-err/dri/releases/latest/download/dri_darwin_arm64.tar.gz
    tar xzf dri_darwin_arm64.tar.gz
    sudo mv dri /usr/local/bin/
    ```

## From Source

Requires Go 1.21+:

```bash
go install github.com/arch-err/dri/cmd/dri@latest
```
