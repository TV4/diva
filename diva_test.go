package diva

import (
	"fmt"
	"net/url"
	"testing"
)

func TestRawURL(t *testing.T) {
	for i, tt := range []struct {
		url  RawURL
		want string
	}{
		{"http://example.com/foo/bar.jpg", "http://example.com/foo/bar.jpg"},
		{"http://diva.cmore.se/image.aspx?id=e4c78001-2854-4151-baa5-a46e070f2cee&formatid=215", "https://img-cdn-cmore.b17g.services/e4c78001-2854-4151-baa5-a46e070f2cee/215.img"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := tt.url.String(); got != tt.want {
				t.Fatalf("url.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	for i, tt := range []struct {
		rawurl string
		want   string
		err    error
	}{
		{"http://example.com/foo/bar.jpg", "", ErrNotDivaURL},
		{"http://diva.cmore.se/image.aspx", "", ErrMissingRequiredArgument},
		{"http://diva.cmore.se/image.aspx?id=e4c78001-2854-4151-baa5-a46e070f2cee&formatid=215", "https://img-cdn-cmore.b17g.services/e4c78001-2854-4151-baa5-a46e070f2cee/215.img", nil},
		{"http://diva.cmore.se/image.aspx?id=b1876803-f5bf-47a6-9a5a-1ef0ee080416&id2=ac4213c2-3d76-4814-80f1-d918700c4eaf&formatid=21", "https://img-cdn-cmore.b17g.services/b1876803-f5bf-47a6-9a5a-1ef0ee080416/ac4213c2-3d76-4814-80f1-d918700c4eaf/21.img", nil},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			u, err := Parse(tt.rawurl)
			if err != nil {
				ue, ok := err.(*url.Error)
				if !ok {
					t.Fatalf("expected to get *url.Error")
				}

				if ue.Err != tt.err {
					t.Fatalf("Parse(%q) returned error: %v\n", tt.rawurl, err)
				}
			}

			if err == nil {
				if got := u.String(); got != tt.want {
					t.Fatalf("url.String() = %q, want %q", got, tt.want)
				}
			}
		})
	}
}

func TestNewURL(t *testing.T) {
	for i, tt := range []struct {
		id1, id2, formatID, want string
	}{
		{"", "", "", ""},
		{"id1", "", "", ""},
		{"id1", "id2", "", ""},
		{"", "id2", "", ""},
		{"", "", "format", ""},
		{"", "id2", "format", ""},
		{"id1", "", "format", "https://img-cdn-cmore.b17g.services/id1/format.img"},
		{"id1", "id2", "format", "https://img-cdn-cmore.b17g.services/id1/id2/format.img"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			u := NewURL(tt.id1, tt.id2, tt.formatID)

			if got := u.String(); got != tt.want {
				t.Fatalf("u.String() = %q, want %q", got, tt.want)
			}
		})
	}
}
