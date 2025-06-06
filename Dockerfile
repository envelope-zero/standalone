FROM golang:1.24.4-alpine as builder

RUN apk add --no-cache make

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG VERSION=0.0.0
RUN CGO_ENABLED=0 make build VERSION=${VERSION}

# Get frontend files from current version
FROM ghcr.io/envelope-zero/frontend:4.2.1 as frontend

# Build final image
FROM scratch
WORKDIR /
COPY --from=frontend /usr/share/nginx/html /public
COPY --from=builder /app/standalone /standalone
ENTRYPOINT ["/standalone"]
