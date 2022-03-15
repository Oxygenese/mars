.PHONY: generate
# generate
wire:
	cd cmd/system && wire

.PHONY: run
run:
	cd cmd/system && go run .


.PHONY: build
# build
build:
	rm -rf ./bin/ && mkdir -p bin/ && go build -o ./bin/ ./...
