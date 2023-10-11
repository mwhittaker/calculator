#!/bin/bash

cloc \
  calc/adder.go calc/calc.go calc/cmd/main.go calc/multiplier.go \
  calc/adder.proto calc/multiplier.proto \
  app.yaml \
  calc/Dockerfile
