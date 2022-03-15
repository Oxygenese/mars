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
	rm -rf ./bin/ && mkdir -p bin/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ ./...
