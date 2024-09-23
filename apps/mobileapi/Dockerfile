FROM golang:1.23.0-alpine AS build_base

ARG FASTTECH_CI_TOKEN

ENV CGO_ENABLED=1
ENV GO111MODULE=on
RUN apk add --no-cache git git gcc g++

# Set the Current Working Directory inside the container
WORKDIR /src

# We want to populate the module cache based on the go.{mod,sum} files.
COPY ../../go.mod .
COPY ../../go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/app ./apps/mobileapi/main.go

# This container exposes port 8080 to the outside world
EXPOSE 5000

# Run the binary program produced by `go install`
ENTRYPOINT ./out/app
