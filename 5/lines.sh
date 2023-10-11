#!/bin/bash

cloc \
  calc/adder.go calc/main.go calc/multiplier.go calc/server.go \
  calc/weaver.toml
