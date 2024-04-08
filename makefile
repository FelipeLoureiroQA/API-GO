.PHONY defaut build test docs cleara

#Variables
APP_NAME=API-GO

#Tasks
defaut: run

run:
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs