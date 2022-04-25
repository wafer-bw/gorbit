lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run
.PHONY: lint

fmt:
	go run golang.org/x/tools/cmd/goimports -w .
.PHONY: fmt

benchmark:
	go test -benchmem -bench . github.com/wafer-bw/gorbit/gravity
.PHONY: benchmark
