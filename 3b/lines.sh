#!/bin/bash

cloc \
  calc/adder.go calc/cmd/main.go calc/server.go multiplier/cmd/main.go multiplier/server.go \
  multiplier/multiplier.proto \
  app.yaml \
  calc/Dockerfile multiplier/Dockerfile
