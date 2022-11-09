#! /bin/sh
rm -rf bin/doup
go build -o bin/doup main.go
chmod +rwx bin/doup
./bin/doup $*