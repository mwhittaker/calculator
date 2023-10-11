#!/usr/bin/env bash

set -euo pipefail

main() {
  cd calc;
  go generate ./...
  go build ./...
  docker build -t themwhittaker/calc:4 .
  docker push themwhittaker/calc:4
}

main "$@"
