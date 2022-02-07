#!/bin/bash

# Move to working dir
cd /accweb

# Remove default config file
rm -rf config.yml

# envsubst to replace set config properly
envsubst < build/docker/docker_config.yml > config.yml

# Run winecfg
winecfg

# Launch accweb main
./main
