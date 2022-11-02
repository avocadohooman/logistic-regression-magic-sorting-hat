format:
	gofmt -d -w .

build:
	go build ./cmd/describe/describe.go
	