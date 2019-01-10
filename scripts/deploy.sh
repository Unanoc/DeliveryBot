#!/bin/bash

docker build --no-cache .. -t vk-bot
docker rm vkbot -f
docker run --name vkbot -d vk-bot /etc/init.d/postgresql start && ./vkbot -config="config.json"
docker ps