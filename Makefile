.PHONY: build run docker

build:
	go build -o bin/server cmd/main.go

run: build
	./bin/server

docker:
	docker build -t $PROJECT_NAME .
