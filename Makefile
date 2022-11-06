format:
	gofmt -d -w .

run: build
	./cmd/describe/describe ../../data/datasets/dataset_train.csv

build:
	go build ./cmd/describe/describe.go