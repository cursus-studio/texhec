FROM golangci/golangci-lint:v2.12.1 AS linter
FROM aquasec/trivy:0.70.0 AS trivy
FROM lycheeverse/lychee:sha-b990100-alpine AS lychee

FROM golang:1.25-bookworm

ENV GOTOOLCHAIN=go1.25.0
ENV PATH="/go/bin:${PATH}"

COPY --from=linter /usr/bin/golangci-lint /usr/bin/golangci-lint
RUN go install github.com/securego/gosec/v2/cmd/gosec@v2.24.6
COPY --from=trivy /usr/local/bin/trivy /usr/bin/trivy
COPY --from=lychee /usr/local/bin/lychee /usr/bin/lychee

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y --no-install-recommends \
    xvfb \
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
    cloc \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/app

USER gouser
