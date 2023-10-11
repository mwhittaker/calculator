#!/usr/bin/env bash

set -euo pipefail

main() {
  for d in calc multiplier; do
    cd $d;
    echo "> Building $d..."
    go generate ./...
    go build ./...
    docker build -t themwhittaker/$d:3 .
    docker push themwhittaker/$d:3
    echo
    cd ..
  done
}

main "$@"
