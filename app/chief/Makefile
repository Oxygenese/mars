VERSION=v1.0.1

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

.PHONY: docker
docker:
	docker build . -t potatozgh/chief:$(VERSION)
	rm -rf ./bin

.PHONY: push
push:
	docker push potatozgh/chief:$(VERSION)
	docker rmi potatozgh/chief:$(VERSION)
.PHONY: register
register:
	make build
	make docker
	make push