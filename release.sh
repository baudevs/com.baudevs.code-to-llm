#!/bin/bash

# Variables
TAG="v0.1.0"
TITLE="v0.1.0 - Initial Release"
NOTES="Initial release of ctllm with support for macOS (Intel & ARM), Linux, and Windows."

# Build and package binaries
echo "Building and packaging binaries..."

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o lib/macos/amd64/ctllm-darwin-amd64 ctllm.go
tar -czvf lib/macos/amd64/ctllm-darwin-amd64.tar.gz -C lib/macos/amd64 ctllm-darwin-amd64

# macOS (ARM)
GOOS=darwin GOARCH=arm64 go build -o lib/macos/arm64/ctllm-darwin-arm64 ctllm.go
tar -czvf lib/macos/arm64/ctllm-darwin-arm64.tar.gz -C lib/macos/arm64 ctllm-darwin-arm64

# Linux (amd64)
GOOS=linux GOARCH=amd64 go build -o lib/linux/amd64/ctllm-linux-amd64 ctllm.go
tar -czvf lib/linux/amd64/ctllm-linux-amd64.tar.gz -C lib/linux/amd64 ctllm-linux-amd64

# Windows (amd64)
GOOS=windows GOARCH=amd64 go build -o lib/windows/amd64/ctllm-windows-amd64.exe ctllm.go
zip lib/windows/amd64/ctllm-windows-amd64.zip -j lib/windows/amd64/ctllm-windows-amd64.exe

echo "Binaries built and packaged."

# Create GitHub Release
echo "Creating GitHub release..."

gh release create $TAG \
  lib/macos/amd64/ctllm-darwin-amd64.tar.gz \
  lib/macos/arm64/ctllm-darwin-arm64.tar.gz \
  lib/linux/amd64/ctllm-linux-amd64.tar.gz \
  lib/windows/amd64/ctllm-windows-amd64.zip \
  --title "$TITLE" \
  --notes "$NOTES"

echo "Release $TAG created successfully."