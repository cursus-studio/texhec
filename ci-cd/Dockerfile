FROM golangci/golangci-lint:v2.12.1 AS linter

FROM golang:1.24-bookworm
COPY --from=linter /usr/bin/golangci-lint /usr/bin/golangci-lint

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

ENV PATH="/go/bin:${PATH}"

WORKDIR /go/src/app
