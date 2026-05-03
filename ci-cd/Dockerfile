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
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.2

ENV PATH="/go/bin:${PATH}"

WORKDIR /go/src/app

ARG CACHE_DATE=1
RUN echo "Last build date: $CACHE_DATE"
