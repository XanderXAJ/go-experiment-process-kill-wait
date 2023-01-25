.PHONY: build build-child build-manager run
.DEFAULT: build

build: build-child build-manager

build-child:
	go build ./cmd/child

build-manager:
	go build ./cmd/manager

run: build
	./manager
