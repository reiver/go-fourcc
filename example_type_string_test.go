package fourcc_test

import (
	"github.com/reiver/go-fourcc"

	"fmt"
)

func ExampleType_String() {

	var colorspace fourcc.Type = fourcc.FOURCC('P','A','L','8')

	s := colorspace.String()

	fmt.Println(s)

	// Output:
	// PAL8
}
