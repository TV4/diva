package diva_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/TV4/diva"
)

func TestConvertRawURL(t *testing.T) {
	for i, tt := range []struct {
		rawurl string
		want   string
	}{
		{"http://example.com/foo/bar.jpg", "http://example.com/foo/bar.jpg"},
		{"http://diva.cmore.se/image.aspx?id=e4c78001-2854-4151-baa5-a46e070f2cee&formatid=215", "https://img-cdn-cmore.b17g.services/e4c78001-2854-4151-baa5-a46e070f2cee/215.img"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := diva.RawURL(tt.rawurl); got != tt.want {
				t.Fatalf("diva.RawURL(%q) = %q, want %q", tt.rawurl, got, tt.want)
			}
		})
	}
}

func ExampleRawURL() {
	fmt.Println(diva.RawURL("http://diva.cmore.se/image.aspx?formatid=221&id=a21630f5-ef51-4632-bf6f-cc94073d3cb1"))

	// Output: https://img-cdn-cmore.b17g.services/a21630f5-ef51-4632-bf6f-cc94073d3cb1/221.img
}

func TestParse(t *testing.T) {
	for i, tt := range []struct {
		rawurl string
		want   string
		err    error
	}{
		{"http://example.com/foo/bar.jpg", "", diva.ErrNotDivaURL},
		{"http://diva.cmore.se/image.aspx", "", diva.ErrMissingRequiredArgument},
		{"http://diva.cmore.se/image.aspx?id=e4c78001-2854-4151-baa5-a46e070f2cee&formatid=215", "https://img-cdn-cmore.b17g.services/e4c78001-2854-4151-baa5-a46e070f2cee/215.img", nil},
		{"http://diva.cmore.se/image.aspx?id=b1876803-f5bf-47a6-9a5a-1ef0ee080416&id2=ac4213c2-3d76-4814-80f1-d918700c4eaf&formatid=21", "https://img-cdn-cmore.b17g.services/b1876803-f5bf-47a6-9a5a-1ef0ee080416/ac4213c2-3d76-4814-80f1-d918700c4eaf/21.img", nil},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			u, err := diva.Parse(tt.rawurl)
			if err != nil {
				ue, ok := err.(*url.Error)
				if !ok {
					t.Fatalf("expected to get *url.Error")
				}

				if ue.Err != tt.err {
					t.Fatalf("diva.Parse(%q) returned error: %v\n", tt.rawurl, err)

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

func ExampleParse() {
	if u, err := diva.Parse("http://diva.cmore.se/image.aspx?formatid=221&id=a21630f5-ef51-4632-bf6f-cc94073d3cb1"); err == nil {
		fmt.Println(u.String())
	}
	// Output: https://img-cdn-cmore.b17g.services/a21630f5-ef51-4632-bf6f-cc94073d3cb1/221.img
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
			u := diva.NewURL(tt.id1, tt.id2, tt.formatID)

			if got := u.String(); got != tt.want {
				t.Fatalf("u.String() = %q, want %q", got, tt.want)
			}
		})
	}
}
