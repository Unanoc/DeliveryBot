#!/bin/bash

psql --command "DROP DATABASE IF EXISTS vkbot;"
psql --command "DROP USER IF EXISTS vkbot;"