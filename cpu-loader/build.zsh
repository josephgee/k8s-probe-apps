#!/bin/zsh

set -euo pipefail

GOOS=linux go build -o probe main.go
docker build -t jgeeatvmware/cpu-loader:1 .
docker push jgeeatvmware/cpu-loader:1
kubectl apply -f pod.yml
