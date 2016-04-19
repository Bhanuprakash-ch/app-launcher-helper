#!/bin/bash

go get -u github.com/go-swagger/go-swagger/cmd/swagger
mkdir -p ./public
swagger generate spec -o ./public/swagger.json
swagger validate ./public/swagger.json
