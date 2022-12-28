travis: test build
	@echo "done"

build:
	go build

test:
	go test -v *_test.go
