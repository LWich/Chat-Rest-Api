# Golang

.PHONY: build-app
build-app:
	go build -o . cmd/api/main.go
start-app:
	./main

# Docker 

run:
	docker run -p 8080:8080 --rm --name chat-rest-api chat-rest-api:test
run-dev:
	docker run -d -p 8080:8080 --rm -v "/home/lwich/Golang/github.com/LWich/chat-rest-api:/app" --name chat-rest-api chat-rest-api:dev
stop:
	docker stop chat-rest-api