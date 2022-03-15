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
	rm -rf ./bin/ && mkdir -p bin/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ ./...

