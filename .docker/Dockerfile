FROM ubuntu:20.10

# Define variables.
ARG GOVERSION=1.17.2
ARG ARCH=amd64

# Download development environment.
RUN apt-get update && \
    apt-get install -y \
    libbpf-dev \
    make \
    clang \
    llvm \
    libelf-dev \
    sudo \
    gcc \
    git

# Install Go specific version.
RUN apt-get install -y wget && \
    wget https://golang.org/dl/go${GOVERSION}.linux-${ARCH}.tar.gz && \
    tar -xf go${GOVERSION}.linux-${ARCH}.tar.gz && \
    mv go/ /usr/local/ && \
    ln -s /usr/local/go/bin/go /usr/local/bin/ && \
    rm -rf go${GOVERSION}.linux-${ARCH}.tar.gz

RUN groupadd --gid 1000 vscode
RUN useradd -s /bin/bash --uid 1000 --gid vscode -m vscode

RUN adduser vscode sudo
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER vscode

# Execute build command.
ENTRYPOINT ["sleep", "infinity"]
