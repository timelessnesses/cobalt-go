build:
	CGO_ENABLED=0 go build -ldflags="-X main.commit=$(shell git rev-parse --short master) -s"