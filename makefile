.PHONY: test
test:
	go test -bench=. -benchmem ./benchmark

.PHONY: run
run:
	go run ./cmd/
