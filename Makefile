travis: test build
	@echo "done"

build:
	go build

test:
	go test -v *_test.go

tidy:
	go mod tidy
	go mod vendor
