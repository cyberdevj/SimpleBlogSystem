#!/bin/bash

unset SBS_ROOT
export SBS_ROOT=$(dirname $(readlink -f $0))

# Grab all golang dependencies
go get ./...

# MongoDB Startup
if !command -v "mongod" &> /dev/null; then
	$SBS_ROOT/bin/mongod --config "$SBS_ROOT/data" &
else
	mongod --dbpath "$SBS_ROOT/data" &
fi

# Build Golang App and Run
make build && $SBS_ROOT/bin/main &

wait
