#!/usr/bin/env bash

set -euo pipefail

main() {
  cd calc;
  echo "> Building calc..."
  go generate ./...
  go build ./...
  docker build -t themwhittaker/calc:1b .
  docker push themwhittaker/calc:1b
}

main "$@"
