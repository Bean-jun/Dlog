.PHONY: all Init-linux-Env Init-win-Env build-linux build-win build docker_build docker_run docker_stop docker_clean run init

WRelease="Release.exe"
LRelease="Release"
ReleaseFolder="Release"
DockerName="dlog"
version=$(shell git log -1 --format=%h)

all: build


Init-linux-Env:
	set CGO_ENABLED=0 GOOS=linux GOARCH=amd64

Init-win-Env:
	set CGO_ENABLED=0 GOOS=windows GOARCH=amd64

build-linux: Init-linux-Env
	go build -o ${ReleaseFolder}/${LRelease}

build-win: Init-win-Env
	go build -o ${ReleaseFolder}/${WRelease}

build: build-linux build-win

docker_build:
	docker build -t dlog:${version} .

docker_run:
	docker run -itd -p 9090:9090 --name ${DockerName} dlog:${version}

docker_stop:
	docker stop ${DockerName}

docker_clean:
	docker rm -f ${DockerName}

# 构建swagger
init:
	swag init

run:
	@go run .
