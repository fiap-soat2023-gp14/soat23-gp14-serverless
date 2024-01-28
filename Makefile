## testing
.PHONY: tests
tests:
	cd src && go test -coverprofile=coverage.out -coverpkg=./... ./... && go tool cover -html=coverage.out
