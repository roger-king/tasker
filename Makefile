DOCKER_TARGET?=ci

setup:
	go mod vendor

app:
	go run example/main.go
	
cleanup:
	rm -rf ./build
	rm -rf ./vendor

default:
	setup