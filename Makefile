DOCKER_TARGET?=ci

setup:
	go mod vendor

dev:
	go run example/main.go
	
cleanup:
	rm -rf ./build
	rm -rf ./vendor

default:
	setup