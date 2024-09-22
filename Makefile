.PHONY: help

## help:	List all commands with descriptions
help: Makefile
	@sed -n "s/^##//p" $<

## test: 			Run tests
test:
	go test ./... -v
