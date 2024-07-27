SHELL := /bin/bash
DIR := $(shell pwd)

# run loan_engine
run: main.go
	go run main.go

# do standard compilation
compile: main.go
	[ -f loan_engine ] && rm loan_engine || true
	go build -ldflags "-s -w" -o ./loan_engine main.go
	upx -9 --lzma loan_engine
	chmod +x loan_engine
