#!/bin/zsh

set -euo pipefail

GOOS=linux go build -o cpu-loader main.go
docker build -t jgeeatvmware/cpu-loader:1 .
docker push jgeeatvmware/cpu-loader:1
kubectl apply -f cpu-loader-pod.yml
