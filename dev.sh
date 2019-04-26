#!/bin/bash

# This file is for local development only!
# It configures and starts accweb for local development.

export ACCWEB_LOGLEVEL=debug
export ACCWEB_WATCH_BUILD_JS=true
export ACCWEB_ALLOWED_ORIGINS=*
export ACCWEB_HOST=localhost:8080

go run main.go
