package diva_test

import (
	"fmt"

	"github.com/TV4/diva"
)

func ExampleCDNRawURL() {
	fmt.Println(diva.CDNRawURL("http://diva.cmore.se/image.aspx?formatid=221&id=a21630f5-ef51-4632-bf6f-cc94073d3cb1"))

	// Output: https://img-cdn-cmore.b17g.services/a21630f5-ef51-4632-bf6f-cc94073d3cb1/221.img
}

func ExampleParse() {
	if u, err := diva.Parse("http://diva.cmore.se/image.aspx?formatid=221&id=a21630f5-ef51-4632-bf6f-cc94073d3cb1"); err == nil {
		fmt.Println(u.String())
	}

	// Output: https://img-cdn-cmore.b17g.services/a21630f5-ef51-4632-bf6f-cc94073d3cb1/221.img
}

func ExampleNewURL() {
	u := diva.NewURL("id1", "id2", "format")
	fmt.Println(u.String())

	// Output: https://img-cdn-cmore.b17g.services/id1/id2/format.img
}
