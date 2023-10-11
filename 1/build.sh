#!/usr/bin/env bash

set -euo pipefail

main() {
  for d in adder multiplier calc; do
    cd $d;
    echo "> Building $d..."
    go generate ./...
    go build ./...
    docker build -t themwhittaker/$d:1 .
    docker push themwhittaker/$d:1
    echo
    cd ..
  done
}

main "$@"
