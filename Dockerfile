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

# Create User (Debian/Ubuntu)
ARG USER=user
RUN groupadd -r -g 1001 ${USER} && \
    useradd -r -u 1001 -g 1001 -m -c "${USER} account" -d /home/${USER} -s /bin/bash ${USER}

WORKDIR /home/${USER}

COPY --chown=${USER}:${USER} go.mod go.sum /home/${USER}/
RUN go mod download

COPY --chown=${USER}:${USER} . /home/${USER}
ARG GO_NEXT_TAG_VERSION="unofficial & built by unknown"
RUN go install -ldflags="-s -w -X 'main.version=${GO_NEXT_TAG_VERSION}'" ./...

# Clear the base image entrypoint
ENTRYPOINT [""]
CMD ["/bin/bash"]

# Timezone
ENV TZ=Europe/Zurich
