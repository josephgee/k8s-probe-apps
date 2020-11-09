#!/bin/zsh

set -euo pipefail

VERSION=2

GOOS=linux go build -o probe main.go
docker build -t jgeeatvmware/cpu-loader:${VERSION} .
docker push jgeeatvmware/cpu-loader:${VERSION}
kubectl apply -f pod.yml
