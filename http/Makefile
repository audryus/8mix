.PHONY: build run test
.SILENT: build run
.DEFAULT_GOAL := run

build:
	@go build -o bin/app

run: build 
	@./bin/app

air: version
	@go build -o bin/2dpoint.exe

version:
	@lazyver semver

test:
	@go clean -testcache
	@go run gotest.tools/gotestsum@latest --packages="./tests/..." --format testdox

docker:
	docker build -f ./build/docker/Dockerfile -t audryus/2dpoint-site .