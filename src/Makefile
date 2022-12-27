.PHONY: all Init-linux-Env Init-win-Env build-linux build-win build docker_build docker_run docker_stop docker_clean run init publish

WRelease="Dlog.exe"
LRelease="Dlog"
ReleaseFolder="Release"
DockerName="dlog"
version=$(shell git log -1 --format=%h)

all: build


Init-linux-Env:
	set CGO_ENABLED=0
	set GOOS=linux
	set GOARCH=amd64

Init-win-Env:
	set CGO_ENABLED=0
	set GOOS=windows
	set GOARCH=amd64

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

publish:
	python shell.py publish

run:
	@go run .
