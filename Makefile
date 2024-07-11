.PHONY: test run

test:
	go test ./... -v

run: test
	go run .

all: test run