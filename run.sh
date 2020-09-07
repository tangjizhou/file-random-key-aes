#!/bin/zsh
go build ./src/main.go
chmod +x main
./main -e -p . -y