#!/bin/zsh

docker compose down && docker rmi goteximg-goweb:latest && docker compose up -d;
