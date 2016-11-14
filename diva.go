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
        fmt.Println(diva.RawURL("http://diva.cmore.se/image.aspx?formatid=221&id=a21630f5-ef51-4632-bf6f-cc94073d3cb1"))
    }

*/
package diva

import (
	"fmt"
	"net/url"
	"strings"
)

var (
	// ErrNotDivaURL is returned when trying to parse a
	// rawurl that isn’t from diva.cmore.se
	ErrNotDivaURL = fmt.Errorf("not a diva URL")

	// ErrMissingRequiredArgument is returned if the id
	// or formatid query parameters are missing
	ErrMissingRequiredArgument = fmt.Errorf("missing required argument")
)

// RawURL converts diva rawurl string into CDN rawurl string
type RawURL string

func (u RawURL) String() string {
	rawurl := string(u)

	if url, err := Parse(rawurl); err == nil {
		return url.String()
	}

	return rawurl
}

// Parse diva rawurl into CDN *url.URL
func Parse(rawurl string) (*url.URL, error) {
	if !strings.Contains(rawurl, "diva.cmore.se/image.aspx") {
		return nil, &url.Error{Op: "parse", URL: rawurl, Err: ErrNotDivaURL}
	}

	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	if u.Query().Get("id") == "" || u.Query().Get("formatid") == "" {
		return nil, &url.Error{Op: "parse", URL: rawurl, Err: ErrMissingRequiredArgument}
	}

	return NewURL(u.Query().Get("id"), u.Query().Get("id2"), u.Query().Get("formatid")), nil
}

// NewURL creates a new image URL with the given ids and format
func NewURL(id, id2, formatID string) *url.URL {
	if id == "" || formatID == "" {
		return &url.URL{}
	}

	return &url.URL{
		Scheme: "https",
		Host:   "img-cdn-cmore.b17g.services",
		Path:   path(id, id2, formatID),
	}
}

func path(id, id2, formatID string) string {
	if id2 != "" {
		return fmt.Sprintf("%s/%s/%s.img", id, id2, formatID)
	}

	return fmt.Sprintf("%s/%s.img", id, formatID)
}
