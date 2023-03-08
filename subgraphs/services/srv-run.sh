#!/bin/bash

# setup deps
go mod download

# go build
go build -o srv server/server.go

# run the server
./srv &

# dump pid

echo $! > ./srv.pid
