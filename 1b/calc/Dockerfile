# Build the server from source.
FROM golang:1.21 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY cmd ./cmd/
RUN cd cmd && CGO_ENABLED=0 go build -o /calc

# Run the server.
FROM alpine:latest
WORKDIR /
COPY --from=build /calc /calc
ENTRYPOINT ["/calc"]
