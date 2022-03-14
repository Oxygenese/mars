.PHONY: generate
# generate
wire:
	cd cmd/auth && wire

.PHONY: run
run:
	cd cmd/auth && go run .

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...
