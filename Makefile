all: build

build: driver/main.go
	go build -o tsgen driver/main.go

sample: build
	./tsgen sample/simple.go

sample2: build
	./tsgen sample/fetch.go


.PHONY: build test

