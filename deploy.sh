#!/usr/bin/env bash
# Pasted into Jenkins to build (will eventually be fleshed out to work with a Docker Hub and Amazon AWS)

echo "Stopping running application"
docker stop ftp-microservice
docker rm ftp-microservice

echo "Building container"
docker build -t byuoitav/ftp-microservice .

echo "Starting the new version"
docker run -d --restart=always --name ftp-microservice -p 8002:8002 byuoitav/ftp-microservice:latest
