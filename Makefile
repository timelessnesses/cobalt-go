build:
	CGO_ENABLED=0 go build -ldflags="-X main.commit=$(git rev-parse --short master) -s"