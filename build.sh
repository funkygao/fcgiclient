#! /bin/bash -e

#===========
# build
#===========
BuildID=$(git rev-parse HEAD | cut -c1-7)
go build -v -ldflags "-X main.BuildID $BuildID"

#===========
# show ver
#===========
./fcgiclient -version
