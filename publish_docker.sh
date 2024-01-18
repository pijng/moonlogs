#!/bin/sh

docker build -t moonlogs:latest .
docker tag moonlogs:latest pijng/moonlogs:latest
docker push pijng/moonlogs:latest