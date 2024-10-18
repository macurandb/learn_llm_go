BINARY_NAME=learn_llm_go

.DEFAULT_GOAL := all

.PHONY: all

clean:
	rm -rf build/
	rm -rf mocks/
	find . -name mock_\*.go -type f -delete
	find . -name *.so -type f -delete
	rm -rf vendor/

tidy:
	go mod tidy
	go mod vendor

goFiles = go list ./... | grep -v ../ccl/grammar

fmt:
	go fmt `$(goFiles)`

vet:
	go vet `$(goFiles)`

lint:
	if ! command -v golangci-lint;then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.55.2; \
		./bin/golangci-lint run; \
	else \
		golangci-lint run; \
	fi

build:
	@echo "Building the project..."
	GOOS=linux GOARCH=amd64 go build -o build/$(BINARY_NAME)
	chmod +x build/$(BINARY_NAME)

run: build
	@echo "Running the project..."
	./build/$(BINARY_NAME)
