#!/usr/bin/env bash

set -euo pipefail

main() {
  cd calc;
  go generate ./...
  go build ./...
}

main "$@"
