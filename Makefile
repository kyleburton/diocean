.PHONY: test clean fmt

all: test

diocean: diocean.go
	go build

test: diocean *_test.go
	go test -test.v -coverprofile=coverage.out

clean:
	rm diocean

fmt:
	go fmt .

cover:
	go test -test.v -coverprofile=coverage.out -covermode=count
	go tool cover -html=coverage.out
