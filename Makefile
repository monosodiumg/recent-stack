# get:
# 	go get -v -t -d ./...

# test:
# 	go test -v -race ./...

# test_cover:
# 	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# .PHONY: get test test_cover

GOBIN ?= $$(go env GOPATH)/bin

.PHONY: install-go-test-coverage
install-go-test-coverage:
	go install github.com/vladopajic/go-test-coverage/v2@latest

.PHONY: check-coverage
check-coverage: install-go-test-coverage
	go test ./... -coverprofile=./cover.out -covermode=atomic
	${GOBIN}/go-test-coverage --config=./.testcoverage.yml