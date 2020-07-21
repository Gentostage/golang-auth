.PHONY: build
build:
	go.exe build -v ./cmd/apiserver

.DEFAULT_GOAL := build