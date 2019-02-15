/*

Package diva converts diva.cmore.se URLs into img-cdn-cmore.b17g.services URLs

Installation

Just go get the package:

    go get -u github.com/TV4/diva

Usage

A small usage example

    package main

    import (
        "fmt"

        "github.com/TV4/diva"
    )

    func main() {
        fmt.Println(diva.CDNRawURL("http://diva.cmore.se/image.aspx?formatid=221&id=a21630f5-ef51-4632-bf6f-cc94073d3cb1"))
    }

*/
package diva

import (
	"fmt"
	"net/url"
)

var (
	// ErrMissingBaseURL is returned when trying to generate an URL without having
	// configured a base URL.
	ErrMissingBaseURL = fmt.Errorf("missing base URL")

	// ErrNotDivaURL is returned when trying to parse a
	// rawurl that isn’t from diva.cmore.se
	ErrNotDivaURL = fmt.Errorf("not a diva URL")

	// ErrMissingRequiredArgument is returned if the id
	// or formatid query parameters are missing
	ErrMissingRequiredArgument = fmt.Errorf("missing required argument")
)

// defaultConverter is the converter used when calling package-level functions.
var defaultConverter = NewConverter("https://img-cdn-cmore.b17g.services/")

// CDNRawURL converts diva rawurl string into CDN rawurl string
func CDNRawURL(rawurl string) string {
	return defaultConverter.CDNRawURL(rawurl)
}

// Parse diva rawurl into CDN *url.URL
func Parse(rawurl string) (*url.URL, error) {
	return defaultConverter.Parse(rawurl)
}

// NewURL creates a new image URL with the given ids and format
func NewURL(id, id2, formatID string) (*url.URL, error) {
	return defaultConverter.NewURL(id, id2, formatID)
}
