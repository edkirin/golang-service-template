EXEC=my-service

run:
	@air

.PHONY: build
build:
	@go build -ldflags "-s -w" -o ./build/${EXEC} ./cmd/service/main.go

upgrade-packages:
	@go get -u ./...
