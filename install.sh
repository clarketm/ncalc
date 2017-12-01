#!/bin/sh

OS=$(uname -s)
ARCH=$(uname -m)

BIN_DIR="/usr/local/bin"
SCRIPT='ncalc'

DARWIN='darwin'
LINUX='linux'

AMD64="amd64"
I386="386"
ARM="arm"
ARM64="arm64"

BUILD_DIR="build"

# Check operating system
case "$OS" in
  [Dd]arwin*)
    BUILD_DIR="${BUILD_DIR}/$DARWIN"
  ;;
  [Ll]inux*)
    BUILD_DIR="${BUILD_DIR}/$LINUX"
  ;;
  *)
    echo "unknown operating system: $OS" >&2
    exit 1
  ;;
esac
echo "$OS operating system detected"

# Check machine architecture
case "$ARCH" in
  [Xx]86*|[Xx]86[-_]64*)
    BUILD_DIR="${BUILD_DIR}_${AMD64}"
  ;;
  i386*|386*)
    BUILD_DIR="${BUILD_DIR}_${I386}"
  ;;
  arm*)
    BUILD_DIR="${BUILD_DIR}_${ARM}"
  ;;
  arm[-_]64*|arm64*)
    BUILD_DIR="${BUILD_DIR}_${ARM64}"
  ;;
  *)
    echo "unknown machine architecture: $ARCH" >&2
    exit 1
  ;;
esac
echo "$ARCH machine architecture detected"

# Make directory (if doesn't exist)
if [ ! -d "$BIN_DIR" ]; then
  echo "creating directory $BIN_DIR"
  sudo mkdir -p "$BIN_DIR"
fi

copy_script () {
  echo "copying $BUILD_DIR.tar.gz/$SCRIPT to $BIN_DIR"
  sudo tar -xzf "$BUILD_DIR.tar.gz" -C "$BIN_DIR"
}

# Copy script
if ! type "$SCRIPT" > /dev/null 2>&1; then
  copy_script
else
  read -r -p "$SCRIPT already exists on your PATH. Do you still want to install? [N/y] " REPLY
  if [ "$REPLY" = 'y' ] || [ "$REPLY" = 'Y' ]; then
    copy_script
  else
    echo "setup aborted" >&2
    exit 1
  fi
fi
