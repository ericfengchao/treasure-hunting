vet:
	go vet ./...

build-tracker:
	go build -o ./tracker github.com/ericfengchao/treasure-hunting/cmd/tracker/

build-player:
	go build -o ./player github.com/ericfengchao/treasure-hunting/cmd/server

build-all: build-tracker build-player
