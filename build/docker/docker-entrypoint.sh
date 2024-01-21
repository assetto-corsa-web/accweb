#!/bin/sh

command=$1

if [ $command = "start" ]; then
    envsubst < docker_config.yml > config.yml
    ./accweb
else
    exec "$@"
fi
