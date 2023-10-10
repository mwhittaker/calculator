# Build the server from source.
FROM golang:1.21 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN GOBIN=/ go install github.com/ServiceWeaver/weaver/cmd/weaver@exec
COPY *.go ./
COPY cmd ./cmd/
RUN cd cmd && CGO_ENABLED=0 go build -o calc

# Run the server.
FROM alpine:latest
WORKDIR /
COPY --from=build /app/calc /app/calc
COPY --from=build /weaver /weaver
ENTRYPOINT ["weaver", "exec", "/app/weaver.toml", "/app/topology.toml"]
