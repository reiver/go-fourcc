package fourcc_test

import (
	"github.com/reiver/go-fourcc"

	"fmt"
)

func ExampleFOURCC() {

	var colorspace fourcc.Type = fourcc.FOURCC('V','P','8','0')

	s := colorspace.String()

	fmt.Println(s)

	// Output:
	// VP80
}
