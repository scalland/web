#!/bin/bash
set -e
VERSION=$(cat VERSION)
BUILD_DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
GIT_COMMIT=$(git rev-parse HEAD)
GIT_STATE=$(git diff --quiet && echo clean || echo dirty)
GIT_SUMMARY=$(git describe --tags --always)

LDFLAGS="-X 'main.Version=${VERSION}' -X 'main.BuildDate=${BUILD_DATE}' -X 'main.GitBranch=${GIT_BRANCH}' -X 'main.GitCommit=${GIT_COMMIT}' -X 'main.GitState=${GIT_STATE}' -X 'main.GitSummary=${GIT_SUMMARY}'"

GOOS=${1:-linux} GOARCH=${2:-amd64} go build -ldflags="${LDFLAGS} -s -w" -o "dist/_${VERSION}_${1:-linux}_${2:-amd64}" .
echo "Built  v${VERSION} for ${1:-linux}/${2:-amd64}"
