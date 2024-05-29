# Locate the binary directory
GOBIN=$(go env GOBIN)
if [ -z "$GOBIN" ]; then
  GOBIN=$(go env GOPATH)/bin
fi

# Remove the okgit binary
rm "$GOBIN/okgit"

# Verify uninstallation
if ! command -v okgit &> /dev/null; then
  echo "okgit successfully uninstalled"
else
  echo "Failed to uninstall okgit"
fi
