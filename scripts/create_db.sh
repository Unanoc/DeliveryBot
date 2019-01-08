#!/bin/bash

psql --command "CREATE USER vkbot WITH SUPERUSER PASSWORD 'vkbot';"
createdb -O vkbot vkbot
psql -d vkbot -c "CREATE EXTENSION IF NOT EXISTS citext;"
psql vkbot -f ./sql/init.sql