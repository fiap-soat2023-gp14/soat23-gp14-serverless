## testing
.PHONY: run-tests
run-tests:
	$(test)
	@cd src && go tool cover -html=../coverage/coverage.out -o=../coverage/index.html

.PHONY: test
test: 
	@mkdir -p coverage
	@cd src && go test -coverprofile=../coverage/coverage.out -coverpkg=./... ./... && go tool cover -func=../coverage/coverage.out

.PHONY: sonar-test
sonar-test:
	@mkdir -p coverage
	@cd src && go test -coverprofile=../coverage/coverage.out ./...