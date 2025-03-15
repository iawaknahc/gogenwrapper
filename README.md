# gogenwrapper

This library is heavily inspired by [felixge/httpsnoop](https://github.com/felixge/httpsnoop).

The difference from it is that this library is more generic.
Notably, it can generate wrappers for the interfaces appearing in [database/sql/driver](https://pkg.go.dev/database/sql/driver).

## Usage example

It is NOT expected that your application depends on this library at runtime.
Instead, you should write your own program making use of [./pkg/gogenwrapper](./pkg/gogenwrapper) to generate the wrappers.

The following files should give you some insights on how to write your own generator program.

- [./pkg/gogenwrapper/internal/databasesqldriver/codegen/main.go](./pkg/gogenwrapper/internal/databasesqldriver/codegen/main.go)
- [./pkg/gogenwrapper/internal/databasesqldriver/gogenerate.go](./pkg/gogenwrapper/internal/databasesqldriver/gogenerate.go)
- [./pkg/gogenwrapper/internal/nethttp/codegen/main.go](./pkg/gogenwrapper/internal/nethttp/codegen/main.go)
- [./pkg/gogenwrapper/internal/nethttp/gogenerate.go](./pkg/gogenwrapper/internal/nethttp/gogenerate.go)
