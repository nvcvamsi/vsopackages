#!/bin/sh
set -eu
cd -- "$(dirname "$0")/.."

go vet ./...
GOOS=js GOARCH=wasm go vet ./...

go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck ./...
GOOS=js GOARCH=wasm staticcheck ./...

govulncheck() {
	tmpf=$(mktemp)
	if ! command govulncheck "$@" >"$tmpf" 2>&1; then
		cat "$tmpf"
	fi
}
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
GOOS=js GOARCH=wasm govulncheck ./...

(
  cd ./internal/examples
  go vet ./...
  staticcheck ./...
  govulncheck ./...
)
(
  cd ./internal/thirdparty
  go vet ./...
  staticcheck ./...
  govulncheck ./...
)
