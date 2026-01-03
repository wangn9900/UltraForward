.PHONY: build-ui build-controller build-agent all

all: build-ui build-controller build-agent

build-ui:
	cd web-new && npm install && npm run build

build-controller:
	go build -o bin/ultra-controller ./cmd/controller

build-agent:
	GOOS=linux GOARCH=amd64 go build -o bin/ultra-agent-linux-amd64 ./cmd/agent
	GOOS=linux GOARCH=arm64 go build -o bin/ultra-agent-linux-arm64 ./cmd/agent

clean:
	rm -rf bin/ web-dist/
