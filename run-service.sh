#!/bin/sh
export $(cat ./env/$1/.env | xargs)
go run services/$1/cmd/main.go