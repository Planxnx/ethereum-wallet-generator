#!/bin/sh

set -e

echo "Are you sure to deploy service to Docker Hub"
read -p "Please confirm [Y/N]? " -n 1 -r
echo # (optional) move to a new line

if [[ $REPLY =~ ^[Yy]$ ]]; then
    TAG="latest"
    REPOSITORY="planxthanee/eth-wallet-gen"
    IMAGE="$REPOSITORY:$TAG"
    # HELM="__helm"

    read -p "Do you want to build code in this repo before deploy [Y/N]? " -n 1 -r
    echo # (optional) move to a new line
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        docker build -t $IMAGE .
        docker push $IMAGE
    fi
fi