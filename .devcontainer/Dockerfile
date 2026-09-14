FROM ubuntu:26.04

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates curl git make gcc g++ wget xz-utils unzip \
    libicu-dev libssl-dev zlib1g libatomic1 && \
    rm -rf /var/lib/apt/lists/*

# Share all tool versions with local development instead of maintaining old pins here.
RUN curl -fsSL https://mise.run | sh
ENV PATH="/root/.local/bin:/root/.local/share/mise/shims:${PATH}"
ENV MISE_GLOBAL_CONFIG_FILE=/opt/tools/mise.toml
ENV MISE_TRUSTED_CONFIG_PATHS=/opt/tools/mise.toml
COPY mise.toml /opt/tools/mise.toml
RUN mise install && mise reshim
