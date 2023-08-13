.PHONY: default run build test docs clean run-with-docs
# Variables
APP_NAME=formValidator

# Tasks
default: run-with-docs

run:
	@go run cmd/formValidator/main.go
run-with-docs:
	@swag init
	@go run cmd/formValidator/main.go
build: 
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
