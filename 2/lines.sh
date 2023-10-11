#!/bin/bash

cloc \
  adder/adder.go adder/cmd/main.go adder/server.go calc/cmd/main.go calc/server.go multiplier/cmd/main.go multiplier/server.go \
  adder/adder.proto multiplier/multiplier.proto \
  app.yaml \
  adder/Dockerfile calc/Dockerfile multiplier/Dockerfile
