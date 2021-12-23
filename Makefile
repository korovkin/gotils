travis: test build
	@echo "done"

build:
	go build

test:
	go test *_test.go
