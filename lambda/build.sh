#!/bin/bash
set -e

echo "Building Go binary for linux/amd64..."
GOOS=linux GOARCH=amd64 go build -o ../tf/build/bootstrap  main.go

echo "Packaging into deployment.zip..."
zip -j ../tf/build/deployment.zip ../tf/build/bootstrap 

echo "Done."
