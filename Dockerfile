#[=======================================================================[
# Description : Docker image containing the go-next-tag binary
#]=======================================================================]

# Docker image repository to use. Use `docker.io` for public images.
ARG IMAGE_BASE_REGISTRY

#### ---- Build ---- ####
FROM ${IMAGE_BASE_REGISTRY}golang:1.22.1-alpine3.19 as build

LABEL maintainer=arash.idelchi

# (can use root througout the image since it's a staged build)
# hadolint ignore=DL3002
USER root

# Basic good practices
SHELL ["/bin/ash", "-o", "pipefail", "-c"]

# timezone
RUN apk add --no-cache \
    tzdata==2024a-r0

WORKDIR /work

COPY go.mod go.sum /work/
RUN go mod download

COPY . /work/
ARG GO_NEXT_TAG_VERSION="unofficial & built by unknown"
RUN go install -ldflags="-s -w -X 'main.version=${GO_NEXT_TAG_VERSION}'" ./...

# Create User (Alpine)
ARG USER=user
RUN addgroup -S -g 1001 ${USER} && \
    adduser -S -u 1001 -G ${USER} -h /home/${USER} -s /bin/ash ${USER}

# Timezone
ENV TZ=Europe/Zurich

#### ---- Standalone ---- ####
FROM scratch as standalone

LABEL maintainer=arash.idelchi

# Copy artifacts from the build stage
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /go/bin/go-next-tag /go-next-tag

USER ${USER}
WORKDIR /home/${USER}

# Clear the base image entrypoint
ENTRYPOINT ["/go-next-tag"]
CMD [""]

# Timezone
ENV TZ=Europe/Zurich

#### ---- App ---- ####
FROM ${IMAGE_BASE_REGISTRY}alpine:3.19

LABEL maintainer=arash.idelchi

USER root

# timezone
RUN apk add --no-cache \
    tzdata==2024a-r0

# Copy artifacts from the build stage
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /go/bin/go-next-tag /usr/local/bin/go-next-tag

USER ${USER}
WORKDIR /home/${USER}

# Clear the base image entrypoint
ENTRYPOINT [""]
CMD ["/bin/ash"]

# Timezone
ENV TZ=Europe/Zurich
