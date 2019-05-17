package diva

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"
)

// Converter converts diva URLs into CDN URLs.
type Converter struct {
	baseURL      string
	cometVersion int
}

// NewConverter returns a new converter instance.
func NewConverter(baseurl string) *Converter {
	return &Converter{
		baseURL:      baseurl,
		cometVersion: 5,
	}
}

func (c *Converter) UseComet6URLParsing() {
	c.cometVersion = 6
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

	nu, err := c.NewURL(u.Query().Get("id"), u.Query().Get("id2"), u.Query().Get("formatid"))
	if err != nil {
		return nil, &url.Error{Op: "parse", URL: rawurl, Err: err}
	}

	return nu, nil
}

// NewURL creates a new image URL with the given ids and format
func (c *Converter) NewURL(id, id2, formatID string) (*url.URL, error) {
	if c.baseURL == "" {
		return nil, ErrMissingBaseURL
	}

	if id == "" || formatID == "" {
		return nil, ErrMissingRequiredArgument
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}

	switch c.cometVersion {
	case 5:
		u.Path = path.Join(u.Path, makeComet5Path(id, id2, formatID))
		return u, nil
	case 6:
		u.Path = path.Join(u.Path, makeComet6Path(id, id2, formatID))
		return u, nil
	default:
		return nil, errors.New("invalid Comet version")
	}
}

func makeComet5Path(id, id2, formatID string) string {
	if id2 != "" {
		return fmt.Sprintf("%s/%s/%s.img", id, id2, formatID)
	}

	return fmt.Sprintf("%s/%s.img", id, formatID)
}

func makeComet6Path(id, id2, formatID string) string {
	if id2 != "" {
		return fmt.Sprintf("%s/%s/%s.jpg", id, id2, formatID)
	}

	return fmt.Sprintf("%s/%s.jpg", id, formatID)
}
