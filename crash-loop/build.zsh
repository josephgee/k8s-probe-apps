#!/bin/zsh

set -euo pipefail

GOOS=linux go build -o probe main.go
docker build -t jgeeatvmware/crash-loop:1 .
docker push jgeeatvmware/crash-loop:1
kubectl apply -f pod.yml
