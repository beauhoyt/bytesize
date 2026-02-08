GO_CMD = go
GO_TEST = $(GO_CMD) test
GO_TOOL = $(GO_CMD) tool
GO_TOOL_COVER = $(GO_TOOL) cover

.PHONY: test
test:
	$(GO_TEST) -v -coverprofile=c.out ./...

.PHONY: cover
cover: test
	$(GO_TOOL_COVER) -func=c.out
	$(GO_TOOL_COVER) -html=c.out

.PHONY: fuzz
fuzz:
	$(GO_TEST) -v -fuzz=FuzzParse -fuzztime=30s bytesize_test.go bytesize.go uint128.go
	$(GO_TEST) -v -fuzz=FuzzFormat -fuzztime=30s bytesize_test.go bytesize.go uint128.go

.PHONY: bench
bench:
	@$(GO_TEST) -v -bench=. -count=10 -run=^$ -fuzz=^$ ./...
