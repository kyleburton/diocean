.PHONY: test clean fmt

all: test

diocean: diocean.go
	go build

test: diocean *_test.go
	go test -test.v

clean:
	rm diocean

fmt:
	go fmt .
