package diva

import (
	"fmt"
	"net/url"
	"strings"
)

// Converter converts diva URLs into CDN URLs.
type Converter struct{}

// NewConverter returns a new converter instance.
func NewConverter() *Converter {
	return &Converter{}
}

// CDNRawURL converts diva rawurl string into CDN rawurl string
func (c *Converter) CDNRawURL(rawurl string) string {
	if u, err := c.Parse(rawurl); err == nil {
		return u.String()
	}

	return rawurl
}

// Parse diva rawurl into CDN *url.URL
func (c *Converter) Parse(rawurl string) (*url.URL, error) {
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

	return c.NewURL(u.Query().Get("id"), u.Query().Get("id2"), u.Query().Get("formatid")), nil
}

// NewURL creates a new image URL with the given ids and format
func (c *Converter) NewURL(id, id2, formatID string) *url.URL {
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
