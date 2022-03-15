.PHONY: generate
# generate
wire:
	cd cmd/chief && wire

.PHONY: run
run:
	cd cmd/chief && go run .

.PHONY: build
# build
build:
	rm -rf ./bin/ && mkdir -p bin/ && go build -o ./bin/ ./...
