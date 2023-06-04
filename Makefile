get:
	go get -v -t ./...

test:
	go test -v -race ./...

test_cover:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: get test test_cover