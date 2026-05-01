# TODO
# agent dockerfile should have sdl2, opengl, golang lint and golang
# Use the official Golang image as the base
FROM golang:1.22-bookworm

ENV DEBIAN_FRONTEND=noninteractive

# Install SDL2 and OpenGL dependencies
# We combine these to keep the layer count down
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

# Ensure the Go bin is in the PATH
ENV PATH="/go/bin:${PATH}"

# Setup a non-root user (optional but recommended for Jenkins)
# Jenkins often overrides the user, but having 'go' environment ready is key
WORKDIR /go/src/app

# The CACHE_DATE arg you mentioned in your Jenkinsfile
# helps force a rebuild of the following layers if needed
ARG CACHE_DATE=1
RUN echo "Last build date: $CACHE_DATE"
