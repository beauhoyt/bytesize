GO_CMD = go
GO_TEST = $(GO_CMD) test

.PHONY: test
test:
	$(GO_TEST) -v ./...

.PHONY: fuzz
fuzz:
	$(GO_TEST) -v -fuzz=FuzzParse -fuzztime=30s bytesize_test.go bytesize.go uint128.go
	$(GO_TEST) -v -fuzz=FuzzFormat -fuzztime=30s bytesize_test.go bytesize.go uint128.go

.PHONY: bench
bench:
	@$(GO_TEST) -v -bench=. -count=10 -run=^$ -fuzz=^$ ./...
