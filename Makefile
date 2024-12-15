.PHONY: format lint test benchmark precommit

format:
	@gofmt -l . | tee /dev/stderr | xargs -r -n 1 echo "File not formatted:" && test -z "$$(gofmt -l .)"

lint:
	@golangci-lint run --timeout=5m

test:
	@go test -v ./... -cover

benchmark:
	@go test ./... -bench=.

precommit: format lint test
