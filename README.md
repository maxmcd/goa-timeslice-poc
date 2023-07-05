# GOA Timeslice bug POC

Minimal POC of Goa breaking when dealing with custom type slices.

You'll need to clone the https://github.com/goadesign/goa repo and make sure it's at the path `../../goadesign/goa` relative to this repo.

From there you can run `make build` to reproduce the build failure.

That command will first run generate with the relative repo:

```
go run goa.design/goa/v3/cmd/goa gen github.com/maxmcd/goa-timeslice-poc/design
```

And then builds the generated code to demonstrate the compiler error:

```
go build ./...
```
