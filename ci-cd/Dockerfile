# TODO
# agent dockerfile should have sdl2, opengl, golang lint and golang
FROM golang:1.24-bookworm

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y --no-install-recommends \
    # SDL2 dependencies
    libsdl2-dev \
    libsdl2-image-dev \
    libsdl2-mixer-dev \
    libsdl2-ttf-dev \
    libsdl2-net-dev \
    # OpenGL dependencies
    mesa-utils \
    libgl1-mesa-dev \
    libglu1-mesa-dev \
    freeglut3-dev \
    # Tools for CI
    pkg-config \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*

# Install golangci-lint
# Using the binary installation method recommended for Linux
RUN go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.12.1

ENV PATH="/go/bin:${PATH}"

WORKDIR /go/src/app
