FROM golang:1.23-alpine3.21 AS builder

WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download
COPY . .

RUN go build -o pdf-generator


FROM node:22-alpine3.21

ENV PUPPETEER_SKIP_CHROMIUM_DOWNLOAD=true
ENV PUPPETEER_EXECUTABLE_PATH="/usr/bin/chromium"

RUN apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz

# Install puppeteer-cli
RUN npm install -g puppeteer-cli

ENV XDG_CONFIG_HOME=/tmp/.chromium
ENV XDG_CACHE_HOME=/tmp/.chromium

# Copy binary
COPY --from=builder /build/pdf-generator /pdf-generator
# Copy template files
COPY --from=builder /build/templates /templates
# Copy tmp folder
COPY --from=builder /build/tmp /tmp

ENTRYPOINT ["/pdf-generator"]