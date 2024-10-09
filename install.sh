#!/bin/bash

# Name of the tool
TOOL_NAME="ctllm"

# Determine the install directory
INSTALL_DIR="/usr/local/bin"

# Build the Go binary
echo "Building the $TOOL_NAME executable..."
if go build -o $TOOL_NAME; then
  echo "$TOOL_NAME executable built successfully."
else
  echo "Failed to build $TOOL_NAME. Please ensure Go is installed."
  exit 1
fi

# Move the binary to the install directory
echo "Installing $TOOL_NAME to $INSTALL_DIR..."
if sudo mv $TOOL_NAME $INSTALL_DIR/; then
  echo "$TOOL_NAME installed successfully in $INSTALL_DIR."
else
  echo "Failed to install $TOOL_NAME. Please check your permissions."
  exit 1
fi

# Ensure the executable has execute permissions
sudo chmod +x $INSTALL_DIR/$TOOL_NAME

echo "$TOOL_NAME installation complete!"
echo "You can now run '$TOOL_NAME --help' to get started."