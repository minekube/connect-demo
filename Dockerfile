FROM golang:1.19 AS build

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o demo demo.go

# Final image
FROM alpine:latest
WORKDIR /demo
COPY --from=build /workspace/demo .
COPY --from=build /workspace/gate.yml .
CMD ["./demo", "--config", "gate.yml"]