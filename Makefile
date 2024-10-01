APP_NAME=goctobot

all:
	@go build -o $(APP_NAME) cmd/application/main.go
	@echo "Done"

fclean:
	@rm $(APP_NAME)
	@echo "Removed App"

tests:
	@go test -v ./...

.PHONY: all fclean tests
