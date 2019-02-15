package diva

import (
	"fmt"
	"net/url"
	"testing"
)

func TestConverter_CDNRawURL(t *testing.T) {
	for i, tt := range []struct {
		rawurl string
		want   string
	}{
		{"http://example.com/foo/bar.jpg", "http://example.com/foo/bar.jpg"},
		{"http://diva.cmore.se/image.aspx?id=e4c78001-2854-4151-baa5-a46e070f2cee&formatid=215", "https://img-cdn-cmore.b17g.services/e4c78001-2854-4151-baa5-a46e070f2cee/215.img"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			conv := NewConverter()

			if got := conv.CDNRawURL(tt.rawurl); got != tt.want {
				t.Fatalf("conv.RawURL(%q) = %q, want %q", tt.rawurl, got, tt.want)
			}
		})
	}
}

func ExampleConverter_CDNRawURL() {
	conv := NewConverter()

	fmt.Println(conv.CDNRawURL("http://diva.cmore.se/image.aspx?formatid=221&id=a21630f5-ef51-4632-bf6f-cc94073d3cb1"))

	// Output: https://img-cdn-cmore.b17g.services/a21630f5-ef51-4632-bf6f-cc94073d3cb1/221.img
}

func TestConverter_Parse(t *testing.T) {
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
			conv := NewConverter()

			u, err := conv.Parse(tt.rawurl)
			if err != nil {
				ue, ok := err.(*url.Error)
				if !ok {
					t.Fatalf("expected to get *url.Error")
				}

				if ue.Err != tt.err {
					t.Fatalf("conv.Parse(%q) returned error: %v\n", tt.rawurl, err)
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

func ExampleConverter_Parse() {
	conv := NewConverter()

	if u, err := conv.Parse("http://diva.cmore.se/image.aspx?formatid=221&id=a21630f5-ef51-4632-bf6f-cc94073d3cb1"); err == nil {
		fmt.Println(u.String())
	}

	// Output: https://img-cdn-cmore.b17g.services/a21630f5-ef51-4632-bf6f-cc94073d3cb1/221.img
}

func TestConverter_NewURL(t *testing.T) {
	for i, tt := range []struct {
		id1, id2, formatID, want string
		wantErr                  bool
	}{
		{"", "", "", "", true},
		{"id1", "", "", "", true},
		{"id1", "id2", "", "", true},
		{"", "id2", "", "", true},
		{"", "", "format", "", true},
		{"", "id2", "format", "", true},
		{"id1", "", "format", "https://img-cdn-cmore.b17g.services/id1/format.img", false},
		{"id1", "id2", "format", "https://img-cdn-cmore.b17g.services/id1/id2/format.img", false},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			conv := NewConverter()

			u, err := conv.NewURL(tt.id1, tt.id2, tt.formatID)

			switch {
			case !tt.wantErr && err == nil:
				if got := u.String(); got != tt.want {
					t.Fatalf("u.String() = %q, want %q", got, tt.want)
				}
			case tt.wantErr && err != nil:
				// ok
			case tt.wantErr && err == nil:
				t.Fatalf("[%d] got nil, want error", i)
			case !tt.wantErr && err != nil:
				t.Fatalf("[%d] got error, want nil: %v", i, err)
			}
		})
	}
}
