# DefEnv

[![GoDoc](https://godoc.org/github.com/reinventer/defenv?status.svg)](https://godoc.org/github.com/reinventer/defenv)
[![Build Status](https://travis-ci.org/reinventer/defenv.svg?branch=master)](https://travis-ci.org/reinventer/defenv)

DefEnv is a go package that contains methods for extracting environment variables and converting them to specified types. If there is no environment variable, the default value is used.

There are two types of methods in a package: ordinary and strict. If there is no environment variable, or if a parsing error occurs, ordinary methods return a default value.
```go
value := defenv.Int("WORKER_NUMBER", 8)
```

Unlike ordinary methods, strict methods return an error when parsing fails.
```go
value, err := defenv.IntStrict("WORKER_NUMBER", 8)
```

## Methods:

| Type          | Ordinary | Strict         |
|---------------|----------|----------------|
| bool          | Bool     | BoolStrict     |
| time.Duration | Duration | DurationStrict |
| float64       | Float64  | Float64Strict  |
| int           | Int      | IntStrict      |
| int64         | Int64    | Int64Strict    |
| string        | String   | -              |
| uint          | Uint     | UintStrict     |
| uint64        | Uint64   | Uint64Strict   |

## Docs

See package documentation at <https://godoc.org/github.com/reinventer/defenv> 