make lint:
	golangci-lint run

make lint-fix:
	golangci-lint run --fix

make lint_fmt:
	golangci-lint fmt

