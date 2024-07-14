generate:
	go generate ./...

test:
	go test -v ./... --cover
