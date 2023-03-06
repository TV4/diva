**Deprecated; this package is no longer maintained.**

# diva

[![Build Status](https://travis-ci.org/TV4/diva.svg?branch=master)](https://travis-ci.org/TV4/diva)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/TV4/diva)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://raw.githubusercontent.com/TV4/diva/master/LICENSE)

Convert diva.cmore.se URLs into img-cdn-cmore.b17g.services URLs

## Installation

    go get -u github.com/TV4/diva

## Usage

### Basic

```go
package main

import (
	"fmt"

	"github.com/TV4/diva"
)

func main() {
	rawurl := "http://diva.cmore.se/image.aspx?formatid=221&id=a21630f5-ef51-4632-bf6f-cc94073d3cb1"

	fmt.Println(diva.CDNRawURL(rawurl))
}
```

### Custom base URL
```go
package main

import (
	"fmt"

	"github.com/TV4/diva"
)

func main() {
	rawurl := "http://diva.cmore.se/image.aspx?formatid=221&id=a21630f5-ef51-4632-bf6f-cc94073d3cb1"

	dc := diva.NewConverter("https://example.com/")

	fmt.Println(dc.CDNRawURL(rawurl))
}
```
