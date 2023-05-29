FROM golang:1.20.4-buster as builder

WORKDIR /app

COPY go.* .

RUN go mod download

RUN go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app cmd/main.go

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/main /app/main

# Run the web service on container startup.
CMD ["/app/main"]
