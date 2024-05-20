build: go.mod go.sum .go-version *.go bin
	@echo "Building sentinel-test-plugin binary..."
	@go build -o bin/sentinel-test-plugin

bin:
	@echo "Creating output directory..."
	@mkdir -p ./bin

.PHONY: clean
clean:
	@echo "Cleaning up build artifacts..."
	@rm -rf ./bin
