APP_NAME := 
VERSION := $(shell cat VERSION)
BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null || echo unknown)
GIT_COMMIT := $(shell git rev-parse HEAD 2>/dev/null || echo none)
GIT_STATE := $(shell git diff --quiet 2>/dev/null && echo clean || echo dirty)
GIT_SUMMARY := $(shell git describe --tags --always 2>/dev/null || echo none)

LDFLAGS := -X 'main.Version=$(VERSION)' \
	-X 'main.BuildDate=$(BUILD_DATE)' \
	-X 'main.GitBranch=$(GIT_BRANCH)' \
	-X 'main.GitCommit=$(GIT_COMMIT)' \
	-X 'main.GitState=$(GIT_STATE)' \
	-X 'main.GitSummary=$(GIT_SUMMARY)'

.PHONY: build run serve migrate seed test vet clean dist-linux-amd64 dist-linux-arm64 dist help

build:
	go build -ldflags="$(LDFLAGS)" -o bin/$(APP_NAME) .

run: build
	./bin/$(APP_NAME) serve

serve: run

migrate: build
	./bin/$(APP_NAME) migrate

seed: build
	./bin/$(APP_NAME) seed

test:
	go test ./...

vet:
	go vet ./...

clean:
	rm -rf bin/ dist/

dist-linux-amd64:
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS) -s -w" -o dist/$(APP_NAME)_$(VERSION)_linux_amd64 .
	-upx --best dist/$(APP_NAME)_$(VERSION)_linux_amd64

dist-linux-arm64:
	mkdir -p dist
	GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS) -s -w" -o dist/$(APP_NAME)_$(VERSION)_linux_arm64 .
	-upx --best dist/$(APP_NAME)_$(VERSION)_linux_arm64

dist: dist-linux-amd64 dist-linux-arm64

help:
	@echo "$(APP_NAME) v$(VERSION)"
	@echo ""
	@echo "Targets:"
	@echo "  build              Build the binary"
	@echo "  run/serve          Build and run the server"
	@echo "  migrate            Run database migrations"
	@echo "  seed               Seed the database"
	@echo "  test               Run tests"
	@echo "  vet                Run go vet"
	@echo "  clean              Remove build artifacts"
	@echo "  dist-linux-amd64   Build linux/amd64 release"
	@echo "  dist-linux-arm64   Build linux/arm64 release"
	@echo "  dist               Build all release binaries"
	@echo "  help               Show this help"
