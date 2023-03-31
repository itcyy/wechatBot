.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux main.go

.PHONY: docker
docker:
	docker build . -t wechatbot:latest
