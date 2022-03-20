VERSION=v1.0.1-beat

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
	rm -rf ./bin/ && mkdir -p bin/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ ./...

.PHONY: docker
docker:
	docker build . -t potatozgh/system:$(VERSION)
	rm -rf ./bin

.PHONY: push
push:
	docker push potatozgh/system:$(VERSION)
	docker rmi potatozgh/system:$(VERSION)