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

This is the problematic generated code:
```
		var zero []time.Time
		if body.Timeslice == zero {
			body.Timeslice = []time.Time{}
		}
```

Creates the compiler error:
```
gen/http/service/server/types.go:36:6: invalid operation: body.Timeslice == zero (slice can only be compared to nil)
```

The code that generates that code is here: https://github.com/goadesign/goa/blob/da7f936afdd523f475bd05c76612850d3cf96ab6/codegen/go_transform.go#L278-L292

Not sure how to fix. `if body.Timeslice == nil` would work, but I'm not sure how to check if the type cannot be compared to a zero value, is that just maps and slices?. A more general fix would be `reflect.DeepEqual(body.Timeslice, zero)`, but that might be frowned upon for an ugly use of reflection.

Here is an example fix by checking if the type is a map or slice by looking for a prefix of `[]` or `map[`: https://github.com/goadesign/goa/compare/v3...maxmcd:goa:timeslice-fix?expand=1
