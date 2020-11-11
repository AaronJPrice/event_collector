.PHONY: test

run:
	@go run main/*.go

test:
	@go test ./...

# Run benchmarks WITHOUT running any tests first
bench:
	@go test ./... -run=XXX -bench=.
