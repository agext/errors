# Extended `errors` package for Go (golang)

[![Release](https://img.shields.io/github/release/agext/errors.svg?style=flat)](https://github.com/agext/errors/releases/latest)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/agext/errors) 
[![Build Status](https://travis-ci.org/agext/errors.svg?branch=master&style=flat)](https://travis-ci.org/agext/errors)
[![Coverage Status](https://coveralls.io/repos/github/agext/errors/badge.svg?style=flat)](https://coveralls.io/github/agext/errors)
[![Go Report Card](https://goreportcard.com/badge/github.com/agext/errors?style=flat)](https://goreportcard.com/report/github.com/agext/errors)

This is a drop-in replacement for the standard [Go](http://golang.org) package with the same name, providing all the standard functionality as well as additional features.

## Project Status

v1.0.1 Stable: Guaranteed no breaking changes to the API in future v1.x releases. Probably safe to use in production, though provided on "AS IS" basis.

This package is being actively maintained. If you encounter any problems or have any suggestions for improvement, please [open an issue](https://github.com/agext/errors/issues). Pull requests are welcome.

## Overview

When you need to retain more information about an error message than a single string allows, just substitute this package for the one in the standard library.

The `New` function still accepts a single string as argument, so no code will be broken. Where you need to include additional information, you can provide it to `New` in a `Desc` structure instead of the string, or you can add it to the error message using one of its setter methods.

The additional information can be used for smarter error handling and logging:
- `Level` differentiates between warnings, regular errors, panics, and fatal errors;
- `Code` allows custom classification and prioritizing, by using ranges or bit-level masks;
- `Info` offers a store for arbitrary data and messages, besides the main error `Text`; the special string `"debug.stack"`, if present as an element in the Info slice, is automatically replaced by a stack trace at the point the error message has been created.

## Installation

```
go get github.com/agext/errors
```

## License

Package errors is released under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
