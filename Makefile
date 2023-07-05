


.PHONY: generate build
generate:
	go run goa.design/goa/v3/cmd/goa gen github.com/maxmcd/goa-timeslice-poc/design

build: generate
	go build ./...
