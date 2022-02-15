#!/usr/bin/bash

if [ "$EUID" -ne 0 ]; then
    sudo -v
fi

go get -v ./...
go build

sudo mv ./ipctl /usr/bin/

