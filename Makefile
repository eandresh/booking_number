LINT_VERSION = v1.45.2
.PHONY: code-format-checks
code-format-check:
	@unformatted_files="$$(gofmt -l .)" \
	&& test -z "$$unformatted_files" || ( printf "Unformatted files: \n$${unformatted_files}\nRun make code-format\n"; exit 1 )

lint:
	golangci-lint run --config golangci.yml

lint-install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(LINT_VERSION)
.PHONY: code-format
code-format:
	goimports -l -w .
	gofmt -l -w .

.PHONY: test
test:
	go test ./...

build-image:
	 DOCKER_BUILDKIT=1  docker build  --ssh default  --force-rm -t template --no-cache .

.PHONY: swagger
swag-install:
	go install github.com/swaggo/swag/cmd/swag@v1.8.7

swag-generate: 
	swag init --parseDependency --parseInternal -d "./cmd/httpserver/main"	
