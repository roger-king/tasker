DOCKER_TARGET?=ci

setup:
	go mod vendor
	go generate

dev: wire
	go run example/main.go

web-dev:
	cd www && BROWSER=none yarn start

cleanup:
	rm -rf ./build
	rm -rf ./vendor

test:
	go test -v github.com/roger-king/tasker/...

wire:
	go generate
	
default:
	setup