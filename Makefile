.PHONY: help

## help:	List all commands with descriptions
help: Makefile
	@sed -n "s/^##//p" $<

## run-app: Run the fetch app
run-app:
	@go run main.go


## test: 			Run tests
test:
	go test ./... -v
