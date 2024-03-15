#[=======================================================================[
# Description : Docker image containing various experimental tooling
#]=======================================================================]

# Docker image repository to use. Use `docker.io` for public images.
ARG IMAGE_BASE_REGISTRY

FROM ${IMAGE_BASE_REGISTRY}golang:1.22.0 as build

LABEL maintainer=arash.idelchi

# (can use root througout the image since it's a staged build)
# hadolint ignore=DL3002
USER root

ARG DEBIAN_FRONTEND=noninteractive

# Basic good practices
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

WORKDIR /work

COPY go.mod go.sum /work/
RUN go mod download

COPY . /work/
ARG GO_NEXT_TAG_VERSION="unofficial & built by unknown"
RUN go install -ldflags="-s -w -X 'main.version=${GO_NEXT_TAG_VERSION}'" ./...

FROM ${IMAGE_BASE_REGISTRY}debian:bookworm-slim

LABEL maintainer=arash.idelchi

USER root

ARG DEBIAN_FRONTEND=noninteractive

# Basic tooling
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Create User (Debian/Ubuntu)
ARG USER=user
RUN groupadd -r -g 1001 ${USER} && \
    useradd -r -u 1001 -g 1001 -m -c "${USER} account" -d /home/${USER} -s /bin/bash ${USER}

COPY --from=build /go/bin/go-next-tag /usr/local/bin/go-next-tag

USER ${USER}
WORKDIR /home/${USER}

# Clear the base image entrypoint
ENTRYPOINT [""]
CMD ["/bin/bash"]

# Timezone
ENV TZ=Europe/Zurich
