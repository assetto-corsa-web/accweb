#!/bin/bash

mkdir -p secrets
openssl genpkey -algorithm RSA -out ./secrets/token.private -pkeyopt rsa_keygen_bits:4096
openssl rsa -pubout -in ./secrets/token.private -out ./secrets/token.public
