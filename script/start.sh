#!/usr/bin/env bash

echo "Running the FancyBot Project..."
export SLACK_TOKEN=$(cat /run/secrets/SLACK_TOKEN)
go run main.go

while true; do sleep 1000; done
