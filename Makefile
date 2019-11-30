DOCKER_TARGET?=ci

setup:
	go mod vendor
	go generate

dev:
	go run example/main.go

cleanup:
	rm -rf ./build
	rm -rf ./vendor

test:
	go test -v github.com/roger-king/tasker/...

wire:
	go generate
	
default:
	setup