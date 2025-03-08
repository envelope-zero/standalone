.PHONY: setup-pre-commit-ci
setup-pre-commit-ci:
# renovate: datasource=github-releases depName=golangci/golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.6

.PHONY: setup
setup: setup-pre-commit-ci
	pre-commit install --hook-type commit-msg --hook-type pre-commit
	go install github.com/air-verse/air@latest

.PHONY: devserver
devserver: frontend
	GIN_MODE=debug air

.PHONY: test
test:
	go test ./... -covermode=atomic -coverprofile=coverage.out -count=1

.PHONY: coverage
coverage: test
	go tool cover -html=coverage.out

.PHONY: frontend
frontend:
	podman create --name frontend-extract ghcr.io/envelope-zero/frontend:4.2.0
	find public -mindepth 1 -not -name ignore-this.txt -delete
	podman cp frontend-extract:/usr/share/nginx/html/. public
	podman rm frontend-extract

VERSION ?= $(shell git rev-parse HEAD)
.PHONY: build
build:
	go build -ldflags "-X github.com/envelope-zero/backend/pkg/router.version=7.0.1"
