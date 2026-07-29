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
    # CI/CD shaders
    clang-format \
    glslang-tools \
    && rm -rf /var/lib/apt/lists/*

RUN curl -sSL -o /usr/bin/cloc https://github.com/AlDanial/cloc/releases/download/v2.08/cloc-2.08.pl \
    && chmod +x /usr/bin/cloc

ENV PIXELORAMA_VERSION=v1.1.10
RUN curl -sSL "https://github.com/Orama-Interactive/Pixelorama/releases/download/${PIXELORAMA_VERSION}/Pixelorama-Linux-64bit.tar.gz" -o pixelorama.tar.gz \
    && mkdir -p /opt/pixelorama \
    && tar -xzf pixelorama.tar.gz -C /opt/pixelorama \
    && rm pixelorama.tar.gz \
    && BINARY=$(find /opt/pixelorama -maxdepth 2 -type f \( -name "Pixelorama*" -a ! -name "*.pck" \)) \
    && chmod +x "$BINARY" \
    && ln -s "$BINARY" /usr/local/bin/pixelorama

USER gouser
WORKDIR /go/src/app
