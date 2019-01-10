#!/bin/bash

docker build --no-cache .. -t vk-bot
docker rm vkbot -f
docker run --name vkbot -d vk-bot ./vkbot -config="config.json"
docker ps