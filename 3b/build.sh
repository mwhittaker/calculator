#!/usr/bin/env bash

set -euo pipefail

main() {
  cd multiplier;
  echo "> Building multiplier..."
  go generate ./...
  go build ./...
  docker build -t themwhittaker/multiplier:3b .
  docker push themwhittaker/multiplier:3b
  echo
  cd ..

  cd calc;
  echo "> Building calc..."
  go generate ./...
  go build ./...
  weaver-kube deploy weaver.toml
  echo
  cd ..
}

main "$@"
