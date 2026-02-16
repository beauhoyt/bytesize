GO_CMD = go
GO_TEST = $(GO_CMD) test
GO_TOOL = $(GO_CMD) tool
GO_TOOL_COVER = $(GO_TOOL) cover
GO_TOOL_CYCLO = $(GO_TOOL) gocyclo
CYCLO_THRESHOLD = 15

.PHONY: test
test:
	$(GO_TEST) -v -coverprofile=c.out ./...
	$(GO_TOOL_COVER) -func=c.out

.PHONY: cover
cover: test
	$(GO_TOOL_COVER) -html=c.out

.PHONY: fuzz
fuzz:
	$(GO_TEST) -v -fuzz=FuzzParse -fuzztime=30s ./...
	$(GO_TEST) -v -fuzz=FuzzFormat -fuzztime=30s ./...

.PHONY: bench
bench:
	@$(GO_TEST) -v -bench=. -count=10 -run=^$ -fuzz=^$ ./...

.IGNORE: cyclo
.PHONY: cyclo
cyclo:
	$(GO_TOOL_CYCLO) -avg -over $(CYCLO_THRESHOLD) -ignore "_test|vendor/" .
	@echo "Return code greater than 0 indicates functions with cyclomatic complexity greater than $(CYCLO_THRESHOLD)."

# Benchmarks for GetMultiplierByUnit* with different function designs and 10 iterations each,
# outputting results to separate files in the bench directory.
.PHONY: benchGetMultiplierByUnit
benchGetMultiplierByUnit:
	mkdir -p bench
	$(GO_TEST) -v -bench='BenchmarkGetMultiplierByUnit.*LongDecimal' -count=10 -run=^$ -fuzz=^$  ./... | tee bench/getMultiplierByUnitLongDecimal.txt
	$(GO_TEST) -v -bench='BenchmarkGetMultiplierByUnit.*LongBinary' -count=10 -run=^$ -fuzz=^$  ./... | tee bench/getMultiplierByUnitLongBinary.txt
	$(GO_TEST) -v -bench='BenchmarkGetMultiplierByUnit.*ShortDecimal' -count=10 -run=^$ -fuzz=^$  ./... | tee bench/getMultiplierByUnitShortDecimal.txt
	$(GO_TEST) -v -bench='BenchmarkGetMultiplierByUnit.*ShortBinary' -count=10 -run=^$ -fuzz=^$  ./... | tee bench/getMultiplierByUnitShortBinary.txt
