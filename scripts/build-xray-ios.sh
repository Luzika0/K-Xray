#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CORE="$ROOT/core"
OUT="$ROOT/ios/Frameworks"
mkdir -p "$OUT"

command -v go >/dev/null || { echo "Go 1.27.3+ is required" >&2; exit 1; }
GO_VERSION="$(go env GOVERSION)"
[[ "$GO_VERSION" == go1.27.* ]] || { echo "Expected Go 1.27.x, found $GO_VERSION" >&2; exit 1; }
if ! command -v gomobile >/dev/null; then
  go install golang.org/x/mobile/cmd/gomobile@latest
fi
gomobile init
(
  cd "$CORE"
  go mod download
  gomobile bind -target=ios -o "$OUT/KXrayEngine.framework" ./engine
)
echo "iOS Xray framework created at ios/Frameworks/KXrayEngine.xcframework"
