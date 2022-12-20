setup:
	go mod tidy

test:
	go test .

.PHONY: setup test