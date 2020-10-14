#!/bin/bash

go test -cover -race ./auth
go test -cover -race ./config
go test -cover -race ./pages
