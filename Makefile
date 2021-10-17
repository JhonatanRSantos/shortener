# Let's ensure that we have the right path variable
PATH = $(shell go env GOPATH)/bin:$(shell echo $$PATH)

all:
	@echo "Select a valid option: run, build-docs"

run: build-docs
	@echo "Runnig the Shortener Project"
	@go run *.go

build-docs:
	@echo "Building the docs"
	@go install github.com/swaggo/swag/cmd/swag@latest
	@export PATH && swag init
