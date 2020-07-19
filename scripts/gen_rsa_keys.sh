#!/bin/bash

echo "creating secrets directory"
mkdir -p secrets
echo "generating private key"
openssl genpkey -algorithm RSA -out ./secrets/token.private -pkeyopt rsa_keygen_bits:4096
echo "generating public key"
openssl rsa -pubout -in ./secrets/token.private -out ./secrets/token.public
