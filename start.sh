#!/usr/bin/env bash
go run -ldflags -s main.go --conf=./conf/conf.dev.toml $1
