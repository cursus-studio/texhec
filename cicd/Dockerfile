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
    # Windows Cross-Compilation
    gcc-mingw-w64-x86-64 \
    g++-mingw-w64-x86-64 \
    mingw-w64-common \
    && rm -rf /var/lib/apt/lists/*

# windows cross compilation dependencies
ENV SDL2_VER=2.30.8
ENV SDL2_IMAGE_VER=2.8.2
ENV SDL2_MIXER_VER=2.8.0
RUN mkdir -p /tmp/sdl-mingw && cd /tmp/sdl-mingw \
    # 1. SDL2 Core
    && curl -sSLO "https://github.com/libsdl-org/SDL/releases/download/release-${SDL2_VER}/SDL2-devel-${SDL2_VER}-mingw.tar.gz" \
    && tar -xzf SDL2-devel-${SDL2_VER}-mingw.tar.gz \
    && cp -r SDL2-${SDL2_VER}/x86_64-w64-mingw32/* /usr/x86_64-w64-mingw32/ \
    # 2. SDL2_image
    && curl -sSLO "https://github.com/libsdl-org/SDL_image/releases/download/release-${SDL2_IMAGE_VER}/SDL2_image-devel-${SDL2_IMAGE_VER}-mingw.tar.gz" \
    && tar -xzf SDL2_image-devel-${SDL2_IMAGE_VER}-mingw.tar.gz \
    && cp -r SDL2_image-${SDL2_IMAGE_VER}/x86_64-w64-mingw32/* /usr/x86_64-w64-mingw32/ \
    # 3. SDL2_mixer
    && curl -sSLO "https://github.com/libsdl-org/SDL_mixer/releases/download/release-${SDL2_MIXER_VER}/SDL2_mixer-devel-${SDL2_MIXER_VER}-mingw.tar.gz" \
    && tar -xzf SDL2_mixer-devel-${SDL2_MIXER_VER}-mingw.tar.gz \
    && cp -r SDL2_mixer-${SDL2_MIXER_VER}/x86_64-w64-mingw32/* /usr/x86_64-w64-mingw32/ \
    # Cleanup
    && rm -rf /tmp/sdl-mingw

RUN if [ "$(dpkg --print-architecture)" = "arm64" ]; then \
        apt-get update && apt-get install -y --no-install-recommends wget gnupg ca-certificates && \
        wget https://ryanfortner.github.io/box64-debs/box64.list -O /etc/apt/sources.list.d/box64.list && \
        wget -qO- https://ryanfortner.github.io/box64-debs/KEY.gpg | gpg --dearmor -o /etc/apt/trusted.gpg.d/box64-debs-archive-keyring.gpg && \
        apt-get update && apt-get install -y --no-install-recommends box64 && \
        rm -rf /var/lib/apt/lists/*; \
    fi

RUN curl -sSL -o /usr/bin/cloc https://github.com/AlDanial/cloc/releases/download/v2.08/cloc-2.08.pl \
    && chmod +x /usr/bin/cloc

# GitHub CLI (used by `cicd release` to create GitHub releases/tags)
RUN mkdir -p -m 755 /etc/apt/keyrings \
    && curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg \
        -o /etc/apt/keyrings/githubcli-archive-keyring.gpg \
    && chmod go+r /etc/apt/keyrings/githubcli-archive-keyring.gpg \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" \
        > /etc/apt/sources.list.d/github-cli.list \
    && apt-get update \
    && apt-get install -y --no-install-recommends gh \
    && rm -rf /var/lib/apt/lists/*

ENV PIXELORAMA_VERSION=v1.1.10
RUN curl -sSL "https://github.com/Orama-Interactive/Pixelorama/releases/download/${PIXELORAMA_VERSION}/Pixelorama-Linux-64bit.tar.gz" -o pixelorama.tar.gz \
    && mkdir -p /opt/pixelorama \
    && tar -xzf pixelorama.tar.gz -C /opt/pixelorama \
    && rm pixelorama.tar.gz \
    && BINARY=$(find /opt/pixelorama -maxdepth 2 -type f \( -name "Pixelorama*" -a ! -name "*.pck" \)) \
    && chmod +x "$BINARY" \
    && ln -s "$BINARY" /usr/local/bin/pixelorama

RUN useradd -m -s /bin/bash gouser
USER gouser
RUN git config --global --add safe.directory "*"
RUN git config --global core.autocrlf true
RUN git config --global core.safecrlf false
WORKDIR /go/src/app
