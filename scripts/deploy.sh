#!/bin/bash

docker build .. -t vk-bot
docker rm vkbot -f
docker run -p 5432:5432 --name vkbot vk-bot 
docker ps