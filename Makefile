all: test

GO_TEST_FLAGS:=-race -count=1 -v
GO_BENCH_PROFILE_FLAGS:=-benchmem -v -cpuprofile=cpuprofile.out -memprofile=memprofile.out -run=^a
GO_BENCH_FLAGS:=-benchmem -v -run=^a
run?=""
runarg?=-run=
component?=.

coverage: GO_TEST_FLAGS:=$(GO_TEST_FLAGS) -coverprofile=coverage.out -cover
coverage: test
	go tool cover -html=coverage.out

bench: GO_TEST_FLAGS:=$(GO_BENCH_FLAGS)
bench: run:=Bench
bench: runarg:=-bench=
bench: test

profile: GO_TEST_FLAGS:=$(GO_BENCH_PROFILE_FLAGS)
profile: run:=Bench
profile: runarg:=-bench=
profile: test

test:
	go test $(GO_TEST_FLAGS) $(component) $(runarg)$(run)

fmt:
	go fmt $(component)

lint:
	-go vet $(component)
	-golint $(component)

deps:
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/tools/cmd/benchcmp
