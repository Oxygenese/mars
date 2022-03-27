VERSION=v1.0.1

.PHONY: generate
# generate
wire:
	cd cmd/cell && wire

.PHONY: run
run:
	cd cmd/cell && go run .

.PHONY: build
# build
build:
	rm -rf ./bin/ && mkdir -p bin/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ ./...

.PHONY: docker
docker:
	docker build . -t potatozgh/cell:$(VERSION)
	rm -rf ./bin

.PHONY: push
push:
	docker push potatozgh/cell:$(VERSION)
	docker rmi potatozgh/cell:$(VERSION)
.PHONY: register
register:
	make wire
	make build
	make docker
	make push