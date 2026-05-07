FROM golangci/golangci-lint:v2.12.1 AS linter
FROM ghcr.io/securego/gosec:v2.24.6 AS gosec
FROM aquasec/trivy:0.25.3 AS trivy

FROM golang:1.24-bookworm
COPY --from=linter /usr/bin/golangci-lint /usr/bin/golangci-lint
COPY --from=gosec /usr/bin/gosec /usr/bin/gosec
COPY --from=trivy /usr/local/bin/trivy /usr/bin/trivy

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
ENV GOTOOLCHAIN=go1.25.0

WORKDIR /go/src/app
